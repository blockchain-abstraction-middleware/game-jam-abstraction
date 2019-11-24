terraform {
  backend "pg" {
    workspaces {
      name = "go-apis"
    }
  }
}

provider "kubernetes" {
}

module "go_api_deployment_game_jam_abstraction" {
  source          = "github.com/blockchain-abstraction-middleware/deployment/modules/deployment"
  namespace       = "go-apis"
  deployment_name = "game-jam-abstraction"
  docker_image    = "bamdockerhub/game-jam-abstraction"
  config_file     = "review.yml"
  config_path     = "/config/"
}