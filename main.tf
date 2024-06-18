# Define a provider (no actual resources are being managed)
provider "null" {}

# A null resource that does nothing
resource "null_resource" "test" {}