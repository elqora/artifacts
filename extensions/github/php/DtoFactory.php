<?php

declare(strict_types=1);

namespace Elqora\Artifact\GitHub;

use Elqora\Artifact\GitHub\Dto\GitHubArtifactReferenceCheckRun;
use Elqora\Artifact\GitHub\Dto\GitHubArtifactReferenceCommit;
use Elqora\Artifact\GitHub\Dto\GitHubArtifactReferenceDeployment;
use Elqora\Artifact\GitHub\Dto\GitHubArtifactReferenceDiscussion;
use Elqora\Artifact\GitHub\Dto\GitHubArtifactReferenceIssue;
use Elqora\Artifact\GitHub\Dto\GitHubArtifactReferencePullRequest;
use Elqora\Artifact\GitHub\Dto\GitHubArtifactReferenceWorkflowRun;
use InvalidArgumentException;

final class DtoFactory
{
    /** @param array<string, mixed> $data */
    public static function reference(array $data): object
    {
        return match ($data['resource'] ?? null) {
            'pull_request' => GitHubArtifactReferencePullRequest::fromArray($data),
            'commit' => GitHubArtifactReferenceCommit::fromArray($data),
            'issue' => GitHubArtifactReferenceIssue::fromArray($data),
            'discussion' => GitHubArtifactReferenceDiscussion::fromArray($data),
            'deployment' => GitHubArtifactReferenceDeployment::fromArray($data),
            'workflow_run' => GitHubArtifactReferenceWorkflowRun::fromArray($data),
            'check_run' => GitHubArtifactReferenceCheckRun::fromArray($data),
            default => throw new InvalidArgumentException('Unsupported GitHub artifact reference discriminator.'),
        };
    }
}
