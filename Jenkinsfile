// Create simple pipeline to React app and deploy on localhost
pipeline {
    //first build Dockerfile
    agent any
    stages {
        stage('Build') {
            steps {
                sh 'docker build -t go-app .'
            }
        }
        stage('Deploy') {
            steps {
                sh 'docker run -d -p 80:80 go-app'
            }
        }
    }
}