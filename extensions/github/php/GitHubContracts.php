<?php

declare(strict_types=1);

namespace Elqora\Artifact\GitHub;

/**
 * @phpstan-type NumberedReference array{repositoryId: positive-int, number: positive-int, nodeId?: non-empty-string}
 * @phpstan-type GitHubArtifactReference (array{resource: 'pull_request'|'issue'|'discussion'}&NumberedReference)|array{resource: 'commit', repositoryId: positive-int, sha: non-empty-string}|array{resource: 'deployment', repositoryId: positive-int, deploymentId: positive-int}|array{resource: 'workflow_run', repositoryId: positive-int, runId: positive-int}|array{resource: 'check_run', repositoryId: positive-int, checkRunId: positive-int}
 * @phpstan-type GitHubArtifactSource array{type: 'provider', provider: 'github', reference: GitHubArtifactReference}
 */
final class GitHubContracts
{
    public const PROVIDER = 'github';
    public const RESOURCES = ['pull_request', 'commit', 'issue', 'discussion', 'deployment', 'workflow_run', 'check_run'];
    private function __construct() {}
}
