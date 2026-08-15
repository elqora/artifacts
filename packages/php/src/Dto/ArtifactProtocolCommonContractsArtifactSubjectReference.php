<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactProtocolCommonContractsArtifactSubjectReference
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('type', $data)) { throw new InvalidArgumentException('Missing required field type for ArtifactProtocolCommonContractsArtifactSubjectReference.'); }
        if (!array_key_exists('id', $data)) { throw new InvalidArgumentException('Missing required field id for ArtifactProtocolCommonContractsArtifactSubjectReference.'); }
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

    public function id(): string
    {
        return $this->data['id'];
    }

    public function hasId(): bool
    {
        return array_key_exists('id', $this->data);
    }

    /** @return array<mixed>|null */
    public function scope(): ?array
    {
        return $this->data['scope'] ?? null;
    }

    public function hasScope(): bool
    {
        return array_key_exists('scope', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['type' => true, 'id' => true, 'scope' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
