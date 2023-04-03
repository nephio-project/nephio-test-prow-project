# nephio-test-prow-project
This repository can be used to test and try the PROW setup. Feel free to submit PRs, test PROW '/' commands like /retest etc.

It also contains sample InRepo Prow config (.prow.yaml) for test jobs that serve as example on how we can use PROW. 

## Prow jobs

Information gathered here is just a quick start to get yourself familiar with managing Prow jobs configured in repository itself.
For detailed information about Prow jobs please refer to the official documentation:

- [ProwJob docs](https://docs.prow.k8s.io/docs/jobs/)
- [Life of a Prow Job](https://docs.prow.k8s.io/docs/life-of-a-prow-job/)
- [Pod Utilities](https://docs.prow.k8s.io/docs/components/pod-utilities/)
- [How to Test a Prow Job](https://docs.prow.k8s.io/docs/build-test-update/#how-to-test-a-prowjob)

## Configuration

Prow job definitions are configured in YAML files. It can be `.prow.yaml` in root of repository or `.prow/` directory containing YAML files. In case both are present `.prow` directory takes precedence. 


This is a sample .prow.yaml:

```
presubmits:
  - name: nmath-test
    decorate: true
    always_run: true
    spec:
      containers:
      - image: golang:1.12.5
        command:
        - /bin/bash
        args:
        - -c
        - "go test"
```

This repo contains `.prow.yaml` that runs simple tests on go code in that repo, feel free to test it out.

## Job types

- **Presubmits** those jobs are being run on pull request content. They are mostly used to validate changes. In the example above job `nmath-test` runs `go test` command every time there's PR submitted against that repository. By default, all presubmit jobs must have success status before PR can be merged. If you set the **optional** parameter to `true`, a PR can be merged even if the job fails.
- **Postsubmits** those jobs are almost the same as the already defined presubmit jobs, they run when you merge the PR so for example trigger Docker image build after feature had been merged.
- **Periodics** jobs run automatically at a scheduled time or every n amount of time. They are not bound to PRs. 

Jobs in both `presubmits` and `postsubmits` categories run in a random order so if in presubmit check you've got lint job and test job on one run they might be in that order in other test job will be executed first. You can check the job status on [`https://prow.nephio.io/`](https://prow.nephio.io/).

For complete config see GoDocs: [`Presubmit`](https://pkg.go.dev/k8s.io/test-infra/prow/config#Presubmit) and [`Postsubmit`](https://pkg.go.dev/k8s.io/test-infra/prow/config#Postsubmit)

## Job images

Example above uses: `golang:1.12` to run `go test ./...` 
Ultimately we'd like to use upstream images wherever's possible, but if there's need for specific changes for example adding `make` etc. this can be done during job run:

```
- name: nmath-lint
  decorate: true
  always_run: true
  spec:
    containers:
    - image: golang:1.12.5
      command:
      - "/bin/sh"
      - "-c"
      - |
        wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.52.2
        bin/golangci-lint run --disable-all -E errcheck -v
```

If more complicated setups are needed those images will have to be built and maintained here with respective Dockerfiles kept in test-infra repository.


## Job naming convention

TODO: 
Naming convention has not yet been established and it's up for discussion, generally it seems like a good idea to use prefix with job type, repository name and type of test so they are distinguishable on PROW UI interface.

## Triggers

Prow presubmit and postsubmit jobs can be triggered based on the following parameters:

- `always_run: true` - run always every time PR is created
- `run_if_changed: {regular expression}` - run if a PR modifies files matching the pattern. If a PR does not modify the files on GitHub job is marked as skipped and doesn't block merging.

**always_run** and **run_if_changed** are mutually exclusive. If you do not set one of them, you can only trigger the job manually by adding a comment to a PR, for example `/test presubmit-myrepo-awesome-linter`

## Interact with Prow

Prow allows you to use GitHub comments to interact with it, for example rerun presubmit jobs on PRs.

> **NOTE:** You can rerun only presubmit jobs.

Full list of commands is available on PROW instance UI under `Menu -> Command Help` or under URL: [`https://prow.nephio.io/command-help`](https://prow.nephio.io/command-help)

For example: if you want to trigger your job again you can add one of those comments to your PR:

`/test all` to rerun all tests
`/retest` to only rerun failed tests
`/test {test-name}` or `/retest {test-name}` to only rerun a specific test. For example, run `/test  presubmit-myrepo-awesome-linter`.

Also if you push anything to branch or fork that you've created your PR from, all defined presubmit jobs will run automatically, you don't have to re-trigger them manually.

After you trigger the job, it appears on [`https://prow.nephio.io/`](https://prow.nephio.io/)


## Prow job tester

Prow main configuration file has job tester job that runs every time file `.prow.yaml` or any file in `.prow` directory has PR created against them to verify the validity of that change. You can see the results of those jobs on Prow UI as well.

## Prow job Environment Variables

Prow exposes few environment variables like `JOB_NAME` or `BUILD_ID` that can be used to dynamically name artifacts etc. For more information see:

[`Job Environment Variables`](https://docs.prow.k8s.io/docs/jobs/#job-environment-variables)

## Secrets

Prow jobs can use secrets located in the same namespace within the cluster
where the jobs are executed, by using the [same mechanism of
podspec](https://kubernetes.io/docs/concepts/configuration/secret/#using-a-secret).
