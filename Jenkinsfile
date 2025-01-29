pipeline {
  agent any
  stages {
    stage('test') {
      steps {
        echo 'Hello World'
      }
    }
    stage('git clone') {
      steps {
        sh 'rm -rf Equity'
        sh 'git clone https://github.com/avisubir350/Equity.git'
        sh 'ls'
        sh 'docker-compose -f Equity/project/docker-compose.yml up -d'
      }
    }
  }
}
