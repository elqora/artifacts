<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\GitHub\Dto;

use InvalidArgumentException;

final readonly class GitHubArtifactSource
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('type', $data)) { throw new InvalidArgumentException('Missing required field type for GitHubArtifactSource.'); }
        if (!array_key_exists('provider', $data)) { throw new InvalidArgumentException('Missing required field provider for GitHubArtifactSource.'); }
        if (!array_key_exists('reference', $data)) { throw new InvalidArgumentException('Missing required field reference for GitHubArtifactSource.'); }
        return new self($data);
    }

    public function type(): string
    {
        return $this->data['type'];
    }

    public function hasType(): bool
    {
        return array_key_exists('type', $this->data);
    }

    public function provider(): string
    {
        return $this->data['provider'];
    }

    public function hasProvider(): bool
    {
        return array_key_exists('provider', $this->data);
    }

    /** @return array<mixed> */
    public function reference(): array
    {
        return $this->data['reference'];
    }

    public function hasReference(): bool
    {
        return array_key_exists('reference', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['type' => true, 'provider' => true, 'reference' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
