package schemalint

import (
	"reflect"
	"testing"
)

func TestAnalyzeName(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		wantCanonical string
		wantSpelling  []string
		wantOrder     []string
	}{
		// Fixed spellings.
		{name: "Ipv6Address", wantCanonical: "Ipv6Address"},
		{name: "IPV6Address", wantCanonical: "Ipv6Address", wantSpelling: []string{`"IPV6" -> "Ipv6"`}},
		{name: "IPv6Address", wantCanonical: "Ipv6Address", wantSpelling: []string{`"IPv6" -> "Ipv6"`}},
		{name: "IpV4CidrBlock", wantCanonical: "Ipv4CidrBlock", wantSpelling: []string{`"IpV4" -> "Ipv4"`}},
		{name: "IPAddress", wantCanonical: "IpAddress", wantSpelling: []string{`"IPAddress" -> "IpAddress"`}},
		{name: "PrivateIPAddresses", wantCanonical: "PrivateIpAddresses", wantSpelling: []string{`"IP" -> "Ip"`}},
		{name: "SecurityGroupIDs", wantCanonical: "SecurityGroupIds", wantSpelling: []string{`"IDs" -> "Ids"`}},
		{name: "SecuritygroupId", wantCanonical: "SecurityGroupId", wantSpelling: []string{`"Securitygroup" -> "SecurityGroup"`}},
		{name: "VPCID", wantCanonical: "VpcId", wantSpelling: []string{`"VPC" -> "Vpc"`, `"ID" -> "Id"`}},
		{name: "EIPAddress", wantCanonical: "EipAddress", wantSpelling: []string{`"EIP" -> "Eip"`}},
		{name: "CIDRBlock", wantCanonical: "CidrBlock", wantSpelling: []string{`"CIDR" -> "Cidr"`}},
		{name: "VPCs", wantCanonical: "Vpcs", wantSpelling: []string{`"VPCs" -> "Vpcs"`}},
		{name: "IDCard", wantCanonical: "IdCard", wantSpelling: []string{`"ID" -> "Id"`}},
		{name: "AllowListIPNum", wantCanonical: "AllowListIpNum", wantSpelling: []string{`"IP" -> "Ip"`}},
		{name: "InstanceDesc", wantCanonical: "InstanceDescription", wantSpelling: []string{`"Desc" -> "Description"`}},
		{name: "STATUS", wantCanonical: "Status", wantSpelling: []string{`"STATUS" -> "Status"`}},
		{name: "ProjectName", wantCanonical: "ProjectName"},
		{name: "SecurityGroupIds", wantCanonical: "SecurityGroupIds"},

		// Fixed word order.
		{name: "AddressIpVersion", wantCanonical: "IpAddressVersion", wantOrder: []string{`"AddressIp" -> "IpAddress"`}},
		{name: "EndpointAddressIP", wantCanonical: "EndpointIpAddress", wantSpelling: []string{`"IP" -> "Ip"`}, wantOrder: []string{`"AddressIp" -> "IpAddress"`}},
		{name: "NameProject", wantCanonical: "ProjectName", wantOrder: []string{`"NameProject" -> "ProjectName"`}},
		{name: "GroupSecurityId", wantCanonical: "SecurityGroupId", wantOrder: []string{`"GroupSecurity" -> "SecurityGroup"`}},

		// Out-of-table acronyms and words that merely start with a table term.
		{name: "VIP", wantCanonical: "VIP"},
		{name: "VIPAddress", wantCanonical: "VIPAddress"},
		{name: "IDCName", wantCanonical: "IDCName"},
		{name: "IPTV", wantCanonical: "IPTV"},
		{name: "UUID", wantCanonical: "UUID"},
		{name: "ISP", wantCanonical: "ISP"},
		{name: "DBInstanceId", wantCanonical: "DBInstanceId"},
		{name: "IpsecConfig", wantCanonical: "IpsecConfig"},
		{name: "IdleTimeout", wantCanonical: "IdleTimeout"},
		{name: "Namespace", wantCanonical: "Namespace"},
		{name: "DescribeMode", wantCanonical: "DescribeMode"},

		// Not PascalCase.
		{name: "key", wantCanonical: "Key"},
		{name: "ipAddress", wantCanonical: "IpAddress", wantSpelling: []string{`"ipAddress" -> "IpAddress"`}},
		{name: "instance_name", wantCanonical: "InstanceName"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			got := analyzeName(testCase.name)

			if got.canonical != testCase.wantCanonical {
				t.Errorf("canonical = %q, want %q", got.canonical, testCase.wantCanonical)
			}
			if !reflect.DeepEqual(got.spellingFixes, testCase.wantSpelling) {
				t.Errorf("spellingFixes = %q, want %q", got.spellingFixes, testCase.wantSpelling)
			}
			if !reflect.DeepEqual(got.orderFixes, testCase.wantOrder) {
				t.Errorf("orderFixes = %q, want %q", got.orderFixes, testCase.wantOrder)
			}
		})
	}
}
