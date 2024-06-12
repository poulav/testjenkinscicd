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
            //    sh "chmod +x -R ${env.WORKSPACE}"
               sh 'go version'
            }
        }
        stage('Initialize') {
            steps {
                // Initialize the Go module
                sh 'go mod init poulav.dev/hostname || true' // This command might fail if go.mod already exists, hence the || true
                sh 'go mod tidy'
            }
        }

        stage('Build') {
            steps {
                // Build the Go program
                sh 'go build -o ${GO_BINARY}'
            }
        }

        stage('Test') {
            steps {
                // Run the Go tests
                sh 'go test ./...'
            }
        }

        stage('Run') {
            steps {
                // Run the built Go program
                sh '''set -x
                ./${GO_BINARY}
                sleep 1
                echo $! > .pidfile
                set +x
                echo 'Now...'
                echo 'Visit http://localhost:8000'
                ''' 
                input message: 'Finished using the web site? (Click "Proceed" to continue)'
                sh '''
                set -x
                kill $(cat .pidfile)
                '''
            }
        }
    }     

    post {
        always {
            // Clean up workspace after build
            cleanWs()
        }
    }   
}
