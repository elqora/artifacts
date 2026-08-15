<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Dto;

use InvalidArgumentException;

final readonly class ArtifactSpec
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
        if (!array_key_exists('schemaVersion', $data)) { throw new InvalidArgumentException('Missing required field schemaVersion for ArtifactSpec.'); }
        if (!array_key_exists('id', $data)) { throw new InvalidArgumentException('Missing required field id for ArtifactSpec.'); }
        if (!array_key_exists('key', $data)) { throw new InvalidArgumentException('Missing required field key for ArtifactSpec.'); }
        if (!array_key_exists('name', $data)) { throw new InvalidArgumentException('Missing required field name for ArtifactSpec.'); }
        if (!array_key_exists('version', $data)) { throw new InvalidArgumentException('Missing required field version for ArtifactSpec.'); }
        if (!array_key_exists('kind', $data)) { throw new InvalidArgumentException('Missing required field kind for ArtifactSpec.'); }
        if (!array_key_exists('valueType', $data)) { throw new InvalidArgumentException('Missing required field valueType for ArtifactSpec.'); }
        return new self($data);
    }

    public function schemaVersion(): string
    {
        return $this->data['schemaVersion'];
    }

    public function hasSchemaVersion(): bool
    {
        return array_key_exists('schemaVersion', $this->data);
    }

    public function id(): string
    {
        return $this->data['id'];
    }

    public function hasId(): bool
    {
        return array_key_exists('id', $this->data);
    }

    public function key(): string
    {
        return $this->data['key'];
    }

    public function hasKey(): bool
    {
        return array_key_exists('key', $this->data);
    }

    public function name(): string
    {
        return $this->data['name'];
    }

    public function hasName(): bool
    {
        return array_key_exists('name', $this->data);
    }

    public function description(): ?string
    {
        return $this->data['description'] ?? null;
    }

    public function hasDescription(): bool
    {
        return array_key_exists('description', $this->data);
    }

    public function version(): int
    {
        return $this->data['version'];
    }

    public function hasVersion(): bool
    {
        return array_key_exists('version', $this->data);
    }

    public function kind(): string
    {
        return $this->data['kind'];
    }

    public function hasKind(): bool
    {
        return array_key_exists('kind', $this->data);
    }

    public function valueType(): string
    {
        return $this->data['valueType'];
    }

    public function hasValueType(): bool
    {
        return array_key_exists('valueType', $this->data);
    }

    /** @return array<mixed>|null */
    public function config(): ?array
    {
        return $this->data['config'] ?? null;
    }

    public function hasConfig(): bool
    {
        return array_key_exists('config', $this->data);
    }

    /** @return array<mixed>|null */
    public function provider(): ?array
    {
        return $this->data['provider'] ?? null;
    }

    public function hasProvider(): bool
    {
        return array_key_exists('provider', $this->data);
    }

    /** @return array<mixed>|null */
    public function requirement(): ?array
    {
        return $this->data['requirement'] ?? null;
    }

    public function hasRequirement(): bool
    {
        return array_key_exists('requirement', $this->data);
    }

    /** @return array<mixed>|null */
    public function lifecycle(): ?array
    {
        return $this->data['lifecycle'] ?? null;
    }

    public function hasLifecycle(): bool
    {
        return array_key_exists('lifecycle', $this->data);
    }

    /** @return array<mixed>|null */
    public function access(): ?array
    {
        return $this->data['access'] ?? null;
    }

    public function hasAccess(): bool
    {
        return array_key_exists('access', $this->data);
    }

    /** @return array<mixed>|null */
    public function privacy(): ?array
    {
        return $this->data['privacy'] ?? null;
    }

    public function hasPrivacy(): bool
    {
        return array_key_exists('privacy', $this->data);
    }

    /** @return array<mixed>|null */
    public function validation(): ?array
    {
        return $this->data['validation'] ?? null;
    }

    public function hasValidation(): bool
    {
        return array_key_exists('validation', $this->data);
    }

    /** @return array<mixed>|null */
    public function verification(): ?array
    {
        return $this->data['verification'] ?? null;
    }

    public function hasVerification(): bool
    {
        return array_key_exists('verification', $this->data);
    }

    /** @return array<mixed>|null */
    public function retention(): ?array
    {
        return $this->data['retention'] ?? null;
    }

    public function hasRetention(): bool
    {
        return array_key_exists('retention', $this->data);
    }

    /** @return array<mixed>|null */
    public function presentation(): ?array
    {
        return $this->data['presentation'] ?? null;
    }

    public function hasPresentation(): bool
    {
        return array_key_exists('presentation', $this->data);
    }

    /** @return array<mixed>|null */
    public function metadata(): ?array
    {
        return $this->data['metadata'] ?? null;
    }

    public function hasMetadata(): bool
    {
        return array_key_exists('metadata', $this->data);
    }

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, ['schemaVersion' => true, 'id' => true, 'key' => true, 'name' => true, 'description' => true, 'version' => true, 'kind' => true, 'valueType' => true, 'config' => true, 'provider' => true, 'requirement' => true, 'lifecycle' => true, 'access' => true, 'privacy' => true, 'validation' => true, 'verification' => true, 'retention' => true, 'presentation' => true, 'metadata' => true]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
