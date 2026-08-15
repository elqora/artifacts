<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactProtocolCommonContractsActorReference
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('type', $data)) { throw new InvalidArgumentException('Missing required field type for ArtifactProtocolCommonContractsActorReference.'); }
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

    public function id(): ?string
    {
        return $this->data['id'] ?? null;
    }

    public function hasId(): bool
    {
        return array_key_exists('id', $this->data);
    }

    public function displayName(): ?string
    {
        return $this->data['displayName'] ?? null;
    }

    public function hasDisplayName(): bool
    {
        return array_key_exists('displayName', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['type' => true, 'id' => true, 'displayName' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
