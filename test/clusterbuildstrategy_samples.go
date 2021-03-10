// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

package test

// MinimalBuildahClusterBuildStrategy defines a
// BuildStrategy for Buildah with two steps
// each of them with different container resources
const MinimalBuildahClusterBuildStrategy = `
apiVersion: shipwright.io/v1alpha1
kind: BuildStrategy
metadata:
  name: buildah
spec:
  buildSteps:
    - name: buildah-bud
      image: quay.io/buildah/stable:latest
      workingDir: /workspace/source
      securityContext:
        privileged: true
      command:
        - /usr/bin/buildah
      args:
        - bud
        - --tag=$(build.output.image)
        - --file=$(build.dockerfile)
        - $(build.source.contextDir)
      resources:
        limits:
          cpu: 500m
          memory: 1Gi
        requests:
          cpu: 500m
          memory: 1Gi
      volumeMounts:
        - name: buildah-images
          mountPath: /var/lib/containers/storage
    - name: buildah-push
      image: quay.io/buildah/stable:latest
      securityContext:
        privileged: true
      command:
        - /usr/bin/buildah
      args:
        - push
        - --tls-verify=false
        - docker://$(build.output.image)
      resources:
        limits:
          cpu: 100m
          memory: 65Mi
        requests:
          cpu: 100m
          memory: 65Mi
      volumeMounts:
        - name: buildah-images
          mountPath: /var/lib/containers/storage
`

// ClusterBuildStrategySingleStep defines a
// BuildStrategy for Buildah with a single step
// and container resources
const ClusterBuildStrategySingleStep = `
apiVersion: shipwright.io/v1alpha1
kind: ClusterBuildStrategy
metadata:
  name: buildah
spec:
  buildSteps:
    - name: buildah-bud
      image: quay.io/buildah/stable:latest
      workingDir: /workspace/source
      securityContext:
        privileged: true
      command:
        - /usr/bin/buildah
      args:
        - bud
        - --tag=$(build.output.image)
        - --file=$(build.dockerfile)
        - $(build.source.contextDir)
      resources:
        limits:
          cpu: 500m
          memory: 1Gi
        requests:
          cpu: 250m
          memory: 65Mi
      volumeMounts:
        - name: buildah-images
          mountPath: /var/lib/containers/storage
    - name: buildah-push
      image: quay.io/buildah/stable:latest
      securityContext:
        privileged: true
      command:
        - /usr/bin/buildah
      args:
        - push
        - --tls-verify=false
        - docker://$(build.output.image)
      resources:
        limits:
          cpu: 500m
          memory: 1Gi
        requests:
          cpu: 250m
          memory: 65Mi
      volumeMounts:
        - name: buildah-images
          mountPath: /var/lib/containers/storage
`

// ClusterBuildStrategySingleStepKaniko is a cluster build strategy based on
// Kaniko, which is very close to the actual Kaniko build strategy example in
// the project
const ClusterBuildStrategySingleStepKaniko = `
apiVersion: shipwright.io/v1alpha1
kind: ClusterBuildStrategy
metadata:
  name: kaniko
spec:
  buildSteps:
    - name: step-build-and-push
      image: gcr.io/kaniko-project/executor:v1.5.1
      workingDir: /workspace/source
      securityContext:
        runAsUser: 0
        capabilities:
          add:
            - CHOWN
            - DAC_OVERRIDE
            - FOWNER
            - SETGID
            - SETUID
            - SETFCAP
            - KILL
      env:
        - name: DOCKER_CONFIG
          value: /tekton/home/.docker
        - name: AWS_ACCESS_KEY_ID
          value: NOT_SET
        - name: AWS_SECRET_KEY
          value: NOT_SET
      command:
        - /kaniko/executor
      args:
        - --skip-tls-verify=true
        - --dockerfile=$(build.dockerfile)
        - --context=/workspace/source/$(build.source.contextDir)
        - --destination=$(build.output.image)
        - --oci-layout-path=/workspace/output/image
        - --snapshotMode=redo
        - --push-retry=3
      resources:
        limits:
          cpu: 500m
          memory: 1Gi
        requests:
          cpu: 250m
          memory: 65Mi
`

// ClusterBuildStrategySingleStepKanikoError is a Kaniko based cluster build
// strategy that has a configuration error (misspelled command flag) so that
// it will fail in Tekton
const ClusterBuildStrategySingleStepKanikoError = `
apiVersion: shipwright.io/v1alpha1
kind: ClusterBuildStrategy
metadata:
  name: kaniko
spec:
  buildSteps:
    - name: step-build-and-push
      image: gcr.io/kaniko-project/executor:v1.5.1
      workingDir: /workspace/source
      securityContext:
        runAsUser: 0
        capabilities:
          add:
            - CHOWN
            - DAC_OVERRIDE
            - FOWNER
            - SETGID
            - SETUID
            - SETFCAP
            - KILL
      env:
        - name: DOCKER_CONFIG
          value: /tekton/home/.docker
        - name: AWS_ACCESS_KEY_ID
          value: NOT_SET
        - name: AWS_SECRET_KEY
          value: NOT_SET
      command:
        - /kaniko/executor
      args:
        - --skips-tlss-verifys=true
        - --dockerfile=$(build.dockerfile)
        - --context=/workspace/source/$(build.source.contextDir)
        - --destination=$(build.output.image)
        - --oci-layout-path=/workspace/output/image
        - --snapshotMode=redo
        - --push-retry=3
      resources:
        limits:
          cpu: 500m
          memory: 1Gi
        requests:
          cpu: 250m
          memory: 65Mi
`

// ClusterBuildStrategyNoOp is a strategy that does nothing and has no dependencies
const ClusterBuildStrategyNoOp = `
apiVersion: shipwright.io/v1alpha1
kind: ClusterBuildStrategy
metadata:
  name: noop
spec:
  buildSteps:
  - name: step-no-and-op
    image: alpine:latest
    workingDir: /workspace/source
    securityContext:
      runAsUser: 0
      capabilities:
        add:
        - CHOWN
        - DAC_OVERRIDE
        - FOWNER
        - SETGID
        - SETUID
        - SETFCAP
        - KILL
    env:
    - name: DOCKER_CONFIG
      value: /tekton/home/.docker
    - name: AWS_ACCESS_KEY_ID
      value: NOT_SET
    - name: AWS_SECRET_KEY
      value: NOT_SET
    command:
    - "true"
    resources:
      limits:
        cpu: 250m
        memory: 128Mi
      requests:
        cpu: 250m
        memory: 128Mi
`

// ClusterBuildStrategySleep30s is a strategy that does only sleep 30 seconds
const ClusterBuildStrategySleep30s = `
apiVersion: build.dev/v1alpha1
kind: ClusterBuildStrategy
metadata:
  name: noop
spec:
  buildSteps:
  - name: sleep30
    image: alpine:latest
    command:
    - sleep
    args:
    - "30s"
    resources:
      limits:
        cpu: 250m
        memory: 128Mi
      requests:
        cpu: 250m
        memory: 128Mi
`

// ClusterBuildStrategyWithAnnotations is a cluster build strategy that contains annotations
const ClusterBuildStrategyWithAnnotations = `
apiVersion: shipwright.io/v1alpha1
kind: ClusterBuildStrategy
metadata:
  annotations:
    kubernetes.io/ingress-bandwidth: 1M
    clusterbuildstrategy.shipwright.io/dummy: aValue
    kubectl.kubernetes.io/last-applied-configuration: anotherValue
  name: kaniko
spec:
  buildSteps:
    - name: step-build-and-push
      image: gcr.io/kaniko-project/executor:v1.5.1
      workingDir: /workspace/source
      securityContext:
        runAsUser: 0
        capabilities:
          add:
            - CHOWN
            - DAC_OVERRIDE
            - FOWNER
            - SETGID
            - SETUID
            - SETFCAP
            - KILL
      env:
        - name: DOCKER_CONFIG
          value: /tekton/home/.docker
        - name: AWS_ACCESS_KEY_ID
          value: NOT_SET
        - name: AWS_SECRET_KEY
          value: NOT_SET
      command:
        - /kaniko/executor
      args:
        - --skip-tls-verify=true
        - --dockerfile=$(build.dockerfile)
        - --context=/workspace/source/$(build.source.contextDir)
        - --destination=$(build.output.image)
        - --oci-layout-path=/workspace/output/image
        - --snapshotMode=redo
        - --push-retry=3
      resources:
        limits:
          cpu: 500m
          memory: 1Gi
        requests:
          cpu: 250m
          memory: 65Mi
`
