<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class HostedArtifactSource
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('type', $data)) { throw new InvalidArgumentException('Missing required field type for HostedArtifactSource.'); }
        if (!array_key_exists('recordType', $data)) { throw new InvalidArgumentException('Missing required field recordType for HostedArtifactSource.'); }
        if (!array_key_exists('recordId', $data)) { throw new InvalidArgumentException('Missing required field recordId for HostedArtifactSource.'); }
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

    public function recordType(): string
    {
        return $this->data['recordType'];
    }

    public function hasRecordType(): bool
    {
        return array_key_exists('recordType', $this->data);
    }

    public function recordId(): string
    {
        return $this->data['recordId'];
    }

    public function hasRecordId(): bool
    {
        return array_key_exists('recordId', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['type' => true, 'recordType' => true, 'recordId' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
