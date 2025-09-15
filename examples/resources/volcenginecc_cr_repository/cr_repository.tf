resource "volcenginecc_cr_repository" "RepositoryDemo" {
  registry     = "test"
  namespace    = "default"
  name         = "RepositoryDemo"
  description  = "RepositoryDemo description"
  access_level = "Public"
}