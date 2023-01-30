import os

from django.core.management.base import BaseCommand, CommandError
from django.utils.module_loading import import_string

from ..protobuf.generators import ModelProtoGenerator


class Command(BaseCommand):
    help = "Generates proto."

    def add_arguments(self, parser):
        parser.add_argument(
            '--model', dest='model', type=str, required=True,
            help='请输入一个模型的路径',
        )
        parser.add_argument(
            '--fields', dest='fields', default=None, type=str,
            help='请给出这个模型要操作的字段'
        )
        parser.add_argument(
            '--file', dest='file', default=None, type=str,
            help='请给出这个proto文件要放入的文件路径'
        )

    def handle(self, *args, **options):
        model = import_string(options['model'])
        fields = options['fields'].split(',') if options['fields'] else None
        filepath = options['file']
        if filepath and os.path.exists(filepath):
            raise CommandError('File "%s" already exists.' % filepath)
        if filepath:
            package = os.path.splitext(os.path.basename(filepath))[0]
        else:
            package = None
        generator = ModelProtoGenerator(
            model=model,
            field_names=fields,
            package=package,
        )
        proto = generator.get_proto()
        if filepath:
            with open(filepath, 'w') as f:
                f.write(proto)
        else:
            self.stdout.write(proto)
