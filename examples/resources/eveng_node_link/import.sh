# Point-to-Point (P2P) link without network ID - network ID is auto-discovered from node interfaces
# <lab_path>|<source_node_id>|<source_port>|<target_node_id>|<target_port>
terraform import eveng_node_link.node '/NodeLink.unl|1|Gi0/1|2|Gi0/1'

# Alternatively with explicit network ID:
# <lab_path>|<network_id>|<source_node_id>|<source_port>|<target_node_id>|<target_port>
terraform import eveng_node_link.node '/NodeLink.unl|1|1|Gi0/1|2|Gi0/1'
