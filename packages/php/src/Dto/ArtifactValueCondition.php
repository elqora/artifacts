<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactValueCondition
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('kind', $data)) { throw new InvalidArgumentException('Missing required field kind for ArtifactValueCondition.'); }
        if (!array_key_exists('artifact', $data)) { throw new InvalidArgumentException('Missing required field artifact for ArtifactValueCondition.'); }
        if (!array_key_exists('operator', $data)) { throw new InvalidArgumentException('Missing required field operator for ArtifactValueCondition.'); }
        if (!array_key_exists('value', $data)) { throw new InvalidArgumentException('Missing required field value for ArtifactValueCondition.'); }
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

    public function artifact(): string
    {
        return $this->data['artifact'];
    }

    public function hasArtifact(): bool
    {
        return array_key_exists('artifact', $this->data);
    }

    public function operator(): string
    {
        return $this->data['operator'];
    }

    public function hasOperator(): bool
    {
        return array_key_exists('operator', $this->data);
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
        return array_diff_key($this->data, ['kind' => true, 'artifact' => true, 'operator' => true, 'value' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
