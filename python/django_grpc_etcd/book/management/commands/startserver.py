import errno
import os
import sys
from datetime import datetime

from django.core.management.base import BaseCommand
from django.conf import settings
from django.utils import autoreload
import grpc
from concurrent import futures

from django.utils.module_loading import import_string

from ..utils.etcd_register import EtcdHandleServ
from django_grpc_etcd import settings as grpc_etcd_settings


class Command(BaseCommand):
    help = "this is used to start grpc-etcd server"

    # requires_system_checks = False

    def __init__(self, stdout=None, stderr=None, no_color=False, force_color=False):
        super().__init__(stdout, stderr, no_color, force_color)
        self.development_mode = None
        self.max_workers = None
        self.address = None
        self.config = grpc_etcd_settings.GRPC_ETCD_CONFIG
        self.etcd_config = self.config.get("ETCD", dict())

    def add_arguments(self, parser):
        parser.add_argument(
            'port', nargs="?", default='10003',
            help='请选择一个端口开启服务'
        )

        parser.add_argument(
            '--max-workers', type=int, default=10, dest='max_workers',
            help='请输入最大进程数量'
        )

        parser.add_argument(
            '--dev', action='store_true', default=True, dest='development_mode',
            help=(
                'Run the server in development mode, This tells Django to user'
                'the auto-reloader and run checks'
            )
        )

    def handle(self, *args, **options):
        grpc_port = self.config.get("GrpcPort", None)
        grpc_max_workers = self.config.get("MaxWorkers")
        self.address = "[::]:{}".format(options['port']) if not grpc_port else "[::]:{}".format(grpc_port)
        self.max_workers = options['max_workers'] if not grpc_max_workers else grpc_max_workers
        self.development_mode = options['development_mode']
        self.run(**options)

    def run(self, **options):
        if self.development_mode:
            if hasattr(autoreload, 'run_with_reloader'):
                autoreload.run_with_reloader(self.inner_run, **options)
            else:
                autoreload.main(self.inner_run, None, options)
        else:
            self.stdout.write((
                                  "starting gRPC server at %(address)s \n"
                              ) % {
                                  "address": self.address
                              })
            self._serve()

    def _serve(self):
        grpc_server = grpc.server(futures.ThreadPoolExecutor(max_workers=self.max_workers))
        # 获取对应的服务对象
        server_class = import_string(self.config.get("SERVER_CLASS"))
        # 从服务对象中获取对应的服务启动方法
        method = getattr(server_class, self.config.get("CLASS_METHOD"))
        # 注册到grpc_server中
        method(server_class(), server=grpc_server)
        grpc_server.add_insecure_port(self.address)
        self.start_etcd()
        grpc_server.start()
        grpc_server.wait_for_termination()

    def start_etcd(self):
        etcd_handle_serv = EtcdHandleServ(service_port=self.etcd_config.get("SERVICE_PORT"),
                                          etcd_ip=self.etcd_config.get("ETCD_IP", "127.0.0.1"),
                                          etcd_port=self.etcd_config.get("ETCD_PORT", 2379),
                                          etcd_prefix=self.etcd_config.get("ETCD_PREFIX"),
                                          service_ip=self.etcd_config.get("SERVICE_IP"))
        etcd_handle_serv.register_service()

    def inner_run(self, *args, **options):
        autoreload.raise_last_exception()

        self.stdout.write("Performing system checks...\n\n")
        self.check(display_num_errors=True)
        # Need to check migrations here, so can't use the
        # requires_migrations_check attribute.
        self.check_migrations()
        now = datetime.now().strftime('%B %d, %Y - %X')
        self.stdout.write(now)
        quit_command = 'CTRL-BREAK' if sys.platform == 'win32' else 'CONTROL-C'
        self.stdout.write((
                              "Django version %(version)s, using settings %(settings)r\n"
                              "Starting development gRPC server at %(address)s\n"
                              "Quit the server with %(quit_command)s.\n"
                          ) % {
                              "version": self.get_version(),
                              "settings": settings.SETTINGS_MODULE,
                              "address": self.address,
                              "quit_command": quit_command,
                          })
        try:
            self._serve()
        except OSError as e:
            # Use helpful error messages instead of ugly tracebacks.
            ERRORS = {
                errno.EACCES: "You don't have permission to access that port.",
                errno.EADDRINUSE: "That port is already in use.",
                errno.EADDRNOTAVAIL: "That IP address can't be assigned to.",
            }
            try:
                error_text = ERRORS[e.errno]
            except KeyError:
                error_text = e
            self.stderr.write("Error: %s" % error_text)
            # Need to use an OS exit because sys.exit doesn't work in a thread
            os._exit(1)
        except KeyboardInterrupt:
            sys.exit(0)
