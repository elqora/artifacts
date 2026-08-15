<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\GitHub\Dto;

use InvalidArgumentException;

final readonly class GitHubArtifactReferenceCommit
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('resource', $data)) { throw new InvalidArgumentException('Missing required field resource for GitHubArtifactReferenceCommit.'); }
        if (!array_key_exists('repositoryId', $data)) { throw new InvalidArgumentException('Missing required field repositoryId for GitHubArtifactReferenceCommit.'); }
        if (!array_key_exists('sha', $data)) { throw new InvalidArgumentException('Missing required field sha for GitHubArtifactReferenceCommit.'); }
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

    public function sha(): string
    {
        return $this->data['sha'];
    }

    public function hasSha(): bool
    {
        return array_key_exists('sha', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['resource' => true, 'repositoryId' => true, 'sha' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
