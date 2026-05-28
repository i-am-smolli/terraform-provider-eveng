# Point-to-Point (P2P) link without network ID - network ID is auto-discovered from node interfaces
terraform import eveng_node_link.node '/NodeLink.unl|1|Gi0/1|2|Gi0/1'

# Alternatively with explicit network ID:
# terraform import eveng_node_link.node '/NodeLink.unl|1|1|Gi0/1|2|Gi0/1'
