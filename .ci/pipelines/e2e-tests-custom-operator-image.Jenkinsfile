def failedTests = []
def testScript

pipeline {

    agent {
        label 'linux'
    }

    options {
        timeout(time: 300, unit: 'MINUTES')
    }

    environment {
        VAULT_ADDR = credentials('vault-addr')
        VAULT_ROLE_ID = credentials('vault-role-id')
        VAULT_SECRET_ID = credentials('vault-secret-id')
        GCLOUD_PROJECT = credentials('k8s-operators-gcloud-project')
    }

    stages {
        stage('Checkout from GitHub') {
            steps {
                checkout scm
            }
        }
        stage('Load common scripts') {
            steps {
                script {
                    testScript = load ".ci/common/tests.groovy"
                }
            }
        }
        stage("Run E2E tests") {
            steps {
                sh """
                    cat >.env <<EOF
GCLOUD_PROJECT = $GCLOUD_PROJECT
OPERATOR_IMAGE = $IMAGE
REGISTRY = eu.gcr.io
REPOSITORY = $GCLOUD_PROJECT
IMG_SUFFIX = -ci
SKIP_DOCKER_COMMAND = true
E2E_JSON = true
TEST_LICENSE = /go/src/github.com/elastic/cloud-on-k8s/.ci/test-license.json
GO_TAGS = release
export LICENSE_PUBKEY = /go/src/github.com/elastic/cloud-on-k8s/.ci/license.key
EOF
                    cat >deployer-config.yml <<EOF
id: gke-ci
overrides:
  clusterName: eck-e2e-$BUILD_NUMBER
  vaultInfo:
    address: $VAULT_ADDR
    roleId: $VAULT_ROLE_ID
    secretId: $VAULT_SECRET_ID
  gke:
    gCloudProject: $GCLOUD_PROJECT
EOF
                """
                script {
                    env.SHELL_EXIT_CODE = sh(returnStatus: true, script: 'make -C .ci get-test-license get-elastic-public-key TARGET=ci-e2e ci')

                    sh 'make -C .ci TARGET=e2e-generate-xml ci'
                    junit "e2e-tests.xml"

                    if (env.SHELL_EXIT_CODE != 0) {
                        failedTests = testScript.getListOfFailedTests()
                    }

                    sh 'exit $SHELL_EXIT_CODE'
                }
            }
        }
    }

    post {
        unsuccessful {
            script {
                if (params.SEND_NOTIFICATIONS) {
                    def msg = testScript.generateSlackMessage("E2E tests failed!", env.BUILD_URL, failedTests)

                    slackSend botUser: true,
                        channel: '#cloud-k8s',
                        color: 'danger',
                        message: msg,
                        tokenCredentialId: 'cloud-ci-slack-integration-token'
                }
            }
        }
        cleanup {
            build job: 'cloud-on-k8s-e2e-cleanup',
                parameters: [string(name: 'GKE_CLUSTER', value: "eck-e2e-${BUILD_NUMBER}")],
                wait: false
            cleanWs()
        }
    }

}
