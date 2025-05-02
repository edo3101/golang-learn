pipeline {
    agent any

    tools {
        go 'go-1.23.1'
    }

    environment {
        SONAR_TOKEN = credentials('SONAR_TOKEN')
    }

    stages {
        stage('Run SonarQube Analysis') {
            steps {
                script {
                    withSinarQubeEnv('SONAR_TOKEN'){
                        sh 'usr/local/bin/sonar-scanner
                        -Dsonar.organization=wm1
                        -Dsonar.projectKey=wm1_golang-learn -Dsonar.sources=.
                        -Dsonar.host.url="https://sonarcloud.io'
                    }
                }
            }
        }

        stage ('Build') {
            steps {
                script{
                    Go application to Nexus
                    sh 'go build -o learn'
                }
                archiveArtifacts 'learn'
            }
        }

        stage ('Build Docker Image') {
            steps {
                script {
                    sh 'docker build -t edo3101/golang-learn-1 .'
                }
            }
        }

        stage ('Push Docker Image') {
            steps {
                script {
                    withCredentials([usernamePassword(credentialsId: 'DOCKER_REGISTRY_CREDENTIALS_ID',
                    usernameVariable: 'DOCKER_USERNAME',
                    passwordVariable: 'DOCKER_PASSWORD')]) {
                        sh """
                            echo $DOCKER_PASSWORD | docker login --username
                            $DOCKER_USERNAME --password-stdin
                            docker push edo3101/golang-learn-1
                        """
                    }
                }
            }
        }

        stage ('Terraform Apply') {
            environment {
                AWS_ACCESS_KEY_ID = credentials('AWS_ACCESS_KEY_ID')
                AWS_SECRET_ACCESS_KEY = credentials('AWS_SECRET_ACCESS_KEY')
            }
            steps {
                script {
                    sh '''
                        export AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID}
                        export AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY}
                        cd ./terraform
                        terraform init
                        terraform apply -auto-approve
                    '''
                }
            }
        }

        stage ('Run Ansible Playbook') {
            steps {
                script {
                    sh '''
                        ansible-playbook ansible/deploy-container.yml
                    '''
                }
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}