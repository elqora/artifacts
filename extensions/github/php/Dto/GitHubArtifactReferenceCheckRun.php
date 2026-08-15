<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\GitHub\Dto;

use InvalidArgumentException;

final readonly class GitHubArtifactReferenceCheckRun
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('resource', $data)) { throw new InvalidArgumentException('Missing required field resource for GitHubArtifactReferenceCheckRun.'); }
        if (!array_key_exists('repositoryId', $data)) { throw new InvalidArgumentException('Missing required field repositoryId for GitHubArtifactReferenceCheckRun.'); }
        if (!array_key_exists('checkRunId', $data)) { throw new InvalidArgumentException('Missing required field checkRunId for GitHubArtifactReferenceCheckRun.'); }
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

    public function checkRunId(): int
    {
        return $this->data['checkRunId'];
    }

    public function hasCheckRunId(): bool
    {
        return array_key_exists('checkRunId', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['resource' => true, 'repositoryId' => true, 'checkRunId' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
