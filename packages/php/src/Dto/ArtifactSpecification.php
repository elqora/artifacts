<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactSpecification
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('schema', $data)) { throw new InvalidArgumentException('Missing required field schema for ArtifactSpecification.'); }
        if (!array_key_exists('version', $data)) { throw new InvalidArgumentException('Missing required field version for ArtifactSpecification.'); }
        if (!array_key_exists('value', $data)) { throw new InvalidArgumentException('Missing required field value for ArtifactSpecification.'); }
        return new self($data);
    }

    public function schema(): string
    {
        return $this->data['schema'];
    }

    public function hasSchema(): bool
    {
        return array_key_exists('schema', $this->data);
    }

    public function version(): int
    {
        return $this->data['version'];
    }

    public function hasVersion(): bool
    {
        return array_key_exists('version', $this->data);
    }

    public function value(): mixed
    {
        return $this->data['value'];
    }

    public function hasValue(): bool
    {
        return array_key_exists('value', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['schema' => true, 'version' => true, 'value' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
