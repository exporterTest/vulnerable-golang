# filename: snmp_example.py

from pysnmp.hlapi import CommunityData, SnmpEngine, UdpTransportTarget, ContextData, getCmd, ObjectType, ObjectIdentity

def get_snmp_data():
    # Setting up SNMP v1 (mpModel=0) community data
    community_data = CommunityData('public', mpModel=0)  # SNMP v1

    # Configuring the SNMP engine and target
    snmp_engine = SnmpEngine()
    transport_target = UdpTransportTarget(('localhost', 161))
    context_data = ContextData()

    # Create an SNMP GET command
    errorIndication, errorStatus, errorIndex, varBinds = next(
        getCmd(
            snmp_engine,
            community_data,
            transport_target,
            context_data,
            ObjectType(ObjectIdentity('1.3.6.1.2.1.1.1.0'))  # sysDescr OID
        )
    )

    # Check for errors and print the result
    if errorIndication:
        print(f"Error: {errorIndication}")
    elif errorStatus:
        print(f"Error: {errorStatus.prettyPrint()}")
    else:
        for varBind in varBinds:
            print(f"{varBind[0]} = {varBind[1]}")

if __name__ == "__main__":
    get_snmp_data()
