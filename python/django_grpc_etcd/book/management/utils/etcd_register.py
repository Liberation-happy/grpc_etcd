import etcd3
import json


class EtcdClient(etcd3.Etcd3Client):
    def get_values_by_key(self, key, **kwargs):
        values, _ = self.get(key, **kwargs)
        values_list = []
        if values is not None:
            try:
                values_list = json.loads(values.decode('utf-8'))
                if not isinstance(values_list, list):
                    raise TypeError()
            except:
                raise Exception()

        return values_list

    def put_values_by_key(self, key, values):
        if not isinstance(values, list):
            raise Exception()
        self.put(key, json.dumps(values))


class EtcdHandleServ():
    def __init__(self, service_port, etcd_ip, etcd_port, etcd_prefix, service_ip):
        self.etcd_ip = etcd_ip
        self.etcd_port = etcd_port
        self.etcd_prefix = etcd_prefix
        self.service_ip = service_ip  # 在本机机器作实验使用
        self.endpoint = f'{service_ip}:{service_port}'

    def register_service(self):
        etcd_client = EtcdClient(host=self.etcd_ip, port=self.etcd_port)
        print(self.etcd_prefix)
        etcd_client.put("/{}/{}".format(self.etcd_prefix, self.endpoint), json.dumps({
            "name":self.etcd_prefix,"addr":self.endpoint,"version":"","weight": 0
        }))

    def logout_service(self):
        etcd_client = EtcdClient(host=self.etcd_ip, port=self.etcd_port)
        key_name = f'{self.etcd_prefix}/grpc'
        with etcd_client.lock(key_name):
            value_list = etcd_client.get_values_by_key(key_name)
            if self.endpoint in value_list:
                value_list.remove(self.endpoint)
                etcd_client.put_values_by_key(key_name, value_list)
