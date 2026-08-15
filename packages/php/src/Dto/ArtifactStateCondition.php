<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactStateCondition
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('kind', $data)) { throw new InvalidArgumentException('Missing required field kind for ArtifactStateCondition.'); }
        if (!array_key_exists('namespace', $data)) { throw new InvalidArgumentException('Missing required field namespace for ArtifactStateCondition.'); }
        if (!array_key_exists('in', $data)) { throw new InvalidArgumentException('Missing required field in for ArtifactStateCondition.'); }
        return new self($data);
    }

    public function kind(): string
    {
        return $this->data['kind'];
    }

    public function hasKind(): bool
    {
        return array_key_exists('kind', $this->data);
    }

    public function namespace(): string
    {
        return $this->data['namespace'];
    }

    public function hasNamespace(): bool
    {
        return array_key_exists('namespace', $this->data);
    }

    /** @return array<mixed> */
    public function in(): array
    {
        return $this->data['in'];
    }

    public function hasIn(): bool
    {
        return array_key_exists('in', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['kind' => true, 'namespace' => true, 'in' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
