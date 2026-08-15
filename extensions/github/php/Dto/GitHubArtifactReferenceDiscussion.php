<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\GitHub\Dto;

use InvalidArgumentException;

final readonly class GitHubArtifactReferenceDiscussion
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('repositoryId', $data)) { throw new InvalidArgumentException('Missing required field repositoryId for GitHubArtifactReferenceDiscussion.'); }
        if (!array_key_exists('number', $data)) { throw new InvalidArgumentException('Missing required field number for GitHubArtifactReferenceDiscussion.'); }
        if (!array_key_exists('resource', $data)) { throw new InvalidArgumentException('Missing required field resource for GitHubArtifactReferenceDiscussion.'); }
        return new self($data);
    }

    public function repositoryId(): int
    {
        return $this->data['repositoryId'];
    }

    public function hasRepositoryId(): bool
    {
        return array_key_exists('repositoryId', $this->data);
    }

    public function number(): int
    {
        return $this->data['number'];
    }

    public function hasNumber(): bool
    {
        return array_key_exists('number', $this->data);
    }

    public function nodeId(): ?string
    {
        return $this->data['nodeId'] ?? null;
    }

    public function hasNodeId(): bool
    {
        return array_key_exists('nodeId', $this->data);
    }

    public function resource(): string
    {
        return $this->data['resource'];
    }

    public function hasResource(): bool
    {
        return array_key_exists('resource', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['repositoryId' => true, 'number' => true, 'nodeId' => true, 'resource' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
