<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\GitHub\Dto;

use InvalidArgumentException;

final readonly class GitHubArtifactReferenceDeployment
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('resource', $data)) { throw new InvalidArgumentException('Missing required field resource for GitHubArtifactReferenceDeployment.'); }
        if (!array_key_exists('repositoryId', $data)) { throw new InvalidArgumentException('Missing required field repositoryId for GitHubArtifactReferenceDeployment.'); }
        if (!array_key_exists('deploymentId', $data)) { throw new InvalidArgumentException('Missing required field deploymentId for GitHubArtifactReferenceDeployment.'); }
        return new self($data);
    }

    public function resource(): string
    {
        return $this->data['resource'];
    }

    public function hasResource(): bool
    {
        return array_key_exists('resource', $this->data);
    }

    public function repositoryId(): int
    {
        return $this->data['repositoryId'];
    }

    public function hasRepositoryId(): bool
    {
        return array_key_exists('repositoryId', $this->data);
    }

    public function deploymentId(): int
    {
        return $this->data['deploymentId'];
    }

    public function hasDeploymentId(): bool
    {
        return array_key_exists('deploymentId', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['resource' => true, 'repositoryId' => true, 'deploymentId' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
