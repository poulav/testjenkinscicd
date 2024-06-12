pipeline {
    agent any

    environment {
        GO_VERSION = '1.20'
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
               sh './test.sh'
            }
        }        
}
