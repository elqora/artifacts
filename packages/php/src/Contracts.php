<?php

declare(strict_types=1);

namespace Elqora\Artifact;

/**
 * Static-analysis contracts for the canonical JSON representation.
 *
 * @phpstan-type Metadata array<string, mixed>
 * @phpstan-type ActorReference array{type: non-empty-string, id?: non-empty-string, displayName?: string}
 * @phpstan-type ScopeReference array{type: non-empty-string, id: non-empty-string}
 * @phpstan-type SubjectReference array{type: non-empty-string, id: non-empty-string, scope?: array<string, string>}
 * @phpstan-type Artifact array{schemaVersion: '1.0', id: non-empty-string, scope?: ScopeReference, kind: non-empty-string, valueType: string, title?: string, description?: string, currentVersionId?: non-empty-string, createdBy: ActorReference, createdAt: string, updatedAt: string, archivedAt?: string, metadata?: Metadata}
 * @phpstan-type InlineSource array{type: 'inline', value: mixed, mediaType?: non-empty-string}
 * @phpstan-type LocalSource array{type: 'local', localId: non-empty-string, filename?: string, mediaType?: non-empty-string, size?: int<0, max>, syncState: string, remoteVersionId?: non-empty-string}
 * @phpstan-type ObjectSource array{type: 'object', objectId: non-empty-string, filename?: string, mediaType?: non-empty-string, size?: int<0, max>, storageProvider?: non-empty-string}
 * @phpstan-type UrlSource array{type: 'url', url: non-empty-string, provider?: non-empty-string, mediaType?: non-empty-string}
 * @phpstan-type HostedSource array{type: 'hosted', recordType: non-empty-string, recordId: non-empty-string}
 * @phpstan-type ProviderSource array{type: 'provider', provider: non-empty-string, reference: mixed}
 * @phpstan-type ArtifactSource InlineSource|LocalSource|ObjectSource|UrlSource|HostedSource|ProviderSource
 * @phpstan-type ArtifactIntegrity array{algorithm: 'sha256'|'sha384'|'sha512', hash: non-empty-string, size?: int<0, max>, verifiedAt?: string}
 * @phpstan-type ArtifactVersion array{schemaVersion: '1.0', id: non-empty-string, artifactId: non-empty-string, version: positive-int, source: ArtifactSource, integrity?: ArtifactIntegrity, createdBy: ActorReference, createdAt: string, note?: string, metadata?: Metadata}
 * @phpstan-type ArtifactLink array{schemaVersion: '1.0', id: non-empty-string, artifactId: non-empty-string, artifactVersionId?: non-empty-string, subject: SubjectReference, role: non-empty-string, note?: string, createdBy: ActorReference, createdAt: string, metadata?: Metadata}
 * @phpstan-type Condition array{kind: string, namespace?: non-empty-string, in?: list<non-empty-string>, artifact?: non-empty-string, operator?: string, value?: mixed, conditions?: list<Condition>, condition?: Condition}
 * @phpstan-type ProviderPolicy array{actors: list<non-empty-string>, mode?: 'single'|'any'|'all', delegation?: 'forbidden'|'allowed'}
 * @phpstan-type RequirementPolicy array{mode: 'required'|'optional'|'conditional', condition?: Condition, blocks?: list<non-empty-string>}
 * @phpstan-type LifecyclePolicy array{createAt?: non-empty-string, editableDuring?: list<non-empty-string>, submitDuring?: list<non-empty-string>, lockAt?: non-empty-string, invalidateOn?: list<non-empty-string>, condition?: Condition}
 * @phpstan-type AccessRule array{actors: list<non-empty-string>, condition?: Condition}
 * @phpstan-type AccessPolicy array{read?: list<AccessRule>, write?: list<AccessRule>, submit?: list<AccessRule>, verify?: list<AccessRule>}
 * @phpstan-type PrivacyPolicy array{classification: 'public'|'internal'|'private'|'sensitive'|'restricted', reveal?: list<array{actors: list<non-empty-string>, when?: Condition, representation: 'hidden'|'masked'|'approximate'|'full'}>, masking?: array{strategy: non-empty-string, config?: array<string, mixed>}, encryption?: array{required: bool, level?: non-empty-string, keyScope?: non-empty-string}}
 * @phpstan-type ValidationPolicy array{mode?: 'strict'|'lenient', rules?: list<array{type: non-empty-string, config?: array<string, mixed>}>}
 * @phpstan-type VerificationPolicy array{required: bool, methods?: list<non-empty-string>, actors?: list<non-empty-string>, condition?: Condition}
 * @phpstan-type RetentionPolicy array{policy: 'forever'}|array{policy: 'duration', days: int<0, max>}|array{policy: 'until', date: string}|array{policy: 'host_defined', key: non-empty-string}
 * @phpstan-type PresentationPolicy array{label?: string, helpText?: string, order?: int|float, display?: non-empty-string, config?: array<string, mixed>}
 * @phpstan-type ValueSchema array{valueType: string, minLength?: int<0, max>, maxLength?: int<0, max>, multiline?: bool, pattern?: string, minimum?: int|float|string, maximum?: int|float|string, integer?: bool, multipleOf?: int|float, currencies?: list<string>, minimumMinorUnits?: int, maximumMinorUnits?: int, mode?: string, requireCoordinates?: bool, allowManualEntry?: bool, minFiles?: int<0, max>, maxFiles?: int<0, max>, acceptedMimeTypes?: list<string>, maxSizeBytes?: int<0, max>, requireTimestamp?: bool, requireLocation?: bool, maxDurationSeconds?: int|float, allowedSchemes?: list<string>, allowedHosts?: list<string>, jsonSchema?: array<string, mixed>, providers?: list<string>, resourceTypes?: list<string>, methods?: list<string>, itemSchema?: ValueSchema, minItems?: int<0, max>, maxItems?: int<0, max>, uniqueItems?: bool}
 * @phpstan-type ArtifactSpec array{schemaVersion: '1.0', id: non-empty-string, key: non-empty-string, name: non-empty-string, description?: string, version: positive-int, kind: non-empty-string, valueType: string, config?: ValueSchema, provider?: ProviderPolicy, requirement?: RequirementPolicy, lifecycle?: LifecyclePolicy, access?: AccessPolicy, privacy?: PrivacyPolicy, validation?: ValidationPolicy, verification?: VerificationPolicy, retention?: RetentionPolicy, presentation?: PresentationPolicy, metadata?: Metadata}
 * @phpstan-type ArtifactSpecSnapshot array{schemaVersion: '1.0', sourceSpecId: non-empty-string, sourceVersion: positive-int, key: non-empty-string, name: non-empty-string, description?: string, kind: non-empty-string, valueType: string, config?: ValueSchema, provider?: ProviderPolicy, requirement?: RequirementPolicy, lifecycle?: LifecyclePolicy, access?: AccessPolicy, privacy?: PrivacyPolicy, validation?: ValidationPolicy, verification?: VerificationPolicy, retention?: RetentionPolicy, presentation?: PresentationPolicy, metadata?: Metadata}
 * @phpstan-type ArtifactRequirement array{schemaVersion: '1.0', id: non-empty-string, key?: non-empty-string, allowedKinds?: list<non-empty-string>, allowedValueTypes?: list<string>, minimumCount?: int<0, max>, maximumCount?: int<0, max>, required: bool, specId?: non-empty-string, metadata?: Metadata}
 * @phpstan-type ArtifactSubmission array{schemaVersion: '1.0', id: non-empty-string, artifactId: non-empty-string, artifactVersionId?: non-empty-string, submittedBy: ActorReference, value?: mixed, submittedAt: string, context?: array{latitude?: float, longitude?: float, deviceId?: non-empty-string, ipAddress?: non-empty-string, userAgent?: string}, metadata?: Metadata}
 * @phpstan-type ArtifactVerification array{schemaVersion: '1.0', id: non-empty-string, artifactId: non-empty-string, artifactVersionId?: non-empty-string, submissionId?: non-empty-string, status: 'pending'|'verified'|'rejected'|'waived', method?: non-empty-string, verifiedBy?: ActorReference, reason?: string, createdAt: string, verifiedAt?: string, metadata?: Metadata}
 */
final class Contracts
{
    private function __construct() {}
}
