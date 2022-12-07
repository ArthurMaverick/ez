#============================================================
#========== Customer Gateway
resource "aws_customer_gateway" "customer_gateway" {
    bgp_asn               = 65000
    ip_address            = "172.0.0.1"
    type                  = "ipsec.1"
    tags = {
        Name              = "${var.ClientVpnName}-customer-gateway"
    }
}
#============================================================
#========== VPN Gateway
resource "aws_vpn_gateway" "vpn_gateway" {
    vpc_id                = "${var.vpc_id}"
    tags = {
        Name              = "${var.ClientVpnName}-vpn-gateway"
    }
}
#============================================================
#========== VPN Gateway Attachment
resource "aws_vpn_gateway_attachment" "vpn_attachment" {
    vpc_id                = "${var.vpc_id}"
    vpn_gateway_id        = "${aws_vpn_gateway.vpn_gateway.id}"
}
#============================================================
#========== VPN Gateway Route Propagation
resource "aws_vpn_gateway_route_propagation" "vpn_gateway_rt_propagation" {
    vpn_gateway_id        = "${aws_vpn_gateway.vpn_gateway.id}"
    route_table_id        = "${var.rt_pvt}"
}
#============================================================
#========== VPN Connection
resource "aws_vpn_connection" "vpn_connection" {
    vpn_gateway_id        = "${aws_vpn_gateway.vpn_gateway.id}"
    customer_gateway_id   = "${aws_customer_gateway.customer_gateway.id}"
    type                  = "ipsec.1"
    static_routes_only    = true
    tunnel1_ike_versions = [ "ikev1" ]
    tags = {
    Name            = "${var.ClientVpnName}-vpn-connection"
  }
}
#============================================================
#========== VPN Connection Route
resource "aws_vpn_connection_route" "this" {
  vpn_connection_id       = "${aws_vpn_connection.vpn_connection.id}"
  destination_cidr_block  = "${var.CidrBlock-ConnectionRule}"
}
