pipeline {
    agent any
    tools {
        go 'golang'
    }

    environment {
        GO_BINARY = 'hostname-app'
    }

    stages {
        stage('Checkout') {
            steps {
                // Checkout the source code from the version control system (e.g., Git)
                checkout scm
            }
        }

        stage('Set Up Go') {
            steps {
               sh "chmod +x -R ${env.WORKSPACE}"
               sh 'go version'
            }
        }
    }        
}
