# Zanroo Dynamic Inventory for Ansible


## Descriptions
A Cli for generates an ansible dynamic inventory from Zanroo Config Management API

## TO-DO
- [x] Mock API Inventory (http://5a1bc601c3630f0012b241f4.mockapi.io/api/v1/)
- [x] Connect Mock API
- [ ] Connect Zanroo Config Management API

## How to build
```bash
build/build.sh "v1.0.0"
```

## Output
```bash
./zanroo-inventory --list
```

```json
{
    "databases": {
        "hosts": ["host1.example.com", "host2.example.com"],
        "vars": {
            "var1": true
        }
    },
    "webservers": ["host3.example.com", "host4.example.com"],
    "zanroo": {
        "hosts": ["host5.example.com", "host6.example.com"],
        "vars": {
            "var1": false
        },
        "children": ["webservers", "databases"]
    }
}
```

## Author Information
DevOps Team