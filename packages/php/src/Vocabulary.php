<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact;

final class Vocabulary
{
    public const PROTOCOL_VERSION = '1.0';
    public const ARTIFACT_VALUE_TYPES = [
        'text',
        'number',
        'boolean',
        'currency',
        'date',
        'datetime',
        'time',
        'location',
        'file',
        'image',
        'video',
        'audio',
        'link',
        'structured',
        'reference',
        'signature',
        'collection',
    ];
    public const ARTIFACT_SOURCE_TYPES = [
        'inline',
        'local',
        'object',
        'url',
        'hosted',
        'provider',
    ];
    public const LOCAL_SYNC_STATES = [
        'local_only',
        'pending_upload',
        'uploading',
        'uploaded',
        'failed',
    ];
    public const INTEGRITY_ALGORITHMS = [
        'sha256',
        'sha384',
        'sha512',
    ];
    public const CONDITION_KINDS = [
        'state',
        'actor',
        'artifact_exists',
        'artifact_value',
        'and',
        'or',
        'not',
    ];
    public const VALUE_OPERATORS = [
        'eq',
        'neq',
        'gt',
        'gte',
        'lt',
        'lte',
        'contains',
        'in',
    ];
    public const PRIVACY_CLASSIFICATIONS = [
        'public',
        'internal',
        'private',
        'sensitive',
        'restricted',
    ];
    public const PRIVACY_REPRESENTATIONS = [
        'hidden',
        'masked',
        'approximate',
        'full',
    ];
    public const VERIFICATION_STATUSES = [
        'pending',
        'verified',
        'rejected',
        'waived',
    ];
}
