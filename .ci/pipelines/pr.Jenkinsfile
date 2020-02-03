pipeline {

    agent {
        label 'linux'
    }

    options {
        timeout(time: 45, unit: 'MINUTES')
    }

    environment {
        VAULT_ADDR = credentials('vault-addr')
        VAULT_ROLE_ID = credentials('vault-role-id')
        VAULT_SECRET_ID = credentials('vault-secret-id')
        GCLOUD_PROJECT = credentials('k8s-operators-gcloud-project')
    }

    stages {
        stage('Check if Docker image needs rebuilding') {
            when {
                expression {
                    notOnlyDocs()
                }
            }
            steps {
                sh 'make -C .ci ci-build-image'
            }
        }
        stage('Run checks') {
            when {
                expression {
                    notOnlyDocs()
                }
            }
            steps {
                echo "Running checks for Go sources..."
                sh 'make -C .ci TARGET=ci-check ci'

                echo "Validating Jenkins pipelines..."
                sh 'make -C .ci TARGET=validate-jenkins-pipelines ci'
            }
        }
        stage('Run tests in parallel') {
            parallel {
                stage("Run unit and integration tests") {
                    when {
                        expression {
                            checkout scm
                            notOnlyDocs()
                        }
                    }
                    agent {
                        label 'linux'
                    }
                    steps {
                        script {
                            env.SHELL_EXIT_CODE = sh(returnStatus: true, script: 'make -C .ci TARGET=ci ci')

                            junit "unit-tests.xml"
                            junit "integration-tests.xml"

                            sh 'exit $SHELL_EXIT_CODE'
                        }
                    }
                }
                stage("Run smoke E2E tests") {
                    when {
                        expression {
                            checkout scm
                            notOnlyDocs()
                        }
                    }
                    steps {
                        script {
                            sh '.ci/setenvconfig pr'
                            env.SHELL_EXIT_CODE = sh(returnStatus: true, script: 'make -C .ci TARGET="build-operator-image ci-e2e" ci')

                            sh 'make -C .ci TARGET=e2e-generate-xml ci'
                            junit "e2e-tests.xml"

                            sh 'exit $SHELL_EXIT_CODE'
                        }
                    }
                }
            }
        }
    }

    post {
        always {
            script {
                if (notOnlyDocs()) {
                    googleStorageUpload bucket: "gs://devops-ci-artifacts/jobs/$JOB_NAME/$BUILD_NUMBER",
                        credentialsId: "devops-ci-gcs-plugin",
                        pattern: "*.tgz",
                        sharedPublicly: true,
                        showInline: true
                }
            }
        }
        cleanup {
            script {
                if (notOnlyDocs()) {
                    build job: 'cloud-on-k8s-e2e-cleanup',
                        parameters: [string(name: 'GKE_CLUSTER', value: "eck-pr-${BUILD_NUMBER}")],
                        wait: false
                }
            }
            cleanWs()
        }
    }
}

def notOnlyDocs() {
    // grep succeeds if there is at least one line without docs/
    return sh (
        script: "git diff --name-status HEAD~1 HEAD | grep -v docs/",
    	returnStatus: true
    ) == 0
}
