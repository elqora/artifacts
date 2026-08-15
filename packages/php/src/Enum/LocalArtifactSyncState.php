<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Enum;

enum LocalArtifactSyncState: string
{
    case V_LOCAL_ONLY = 'local_only';
    case V_PENDING_UPLOAD = 'pending_upload';
    case V_UPLOADING = 'uploading';
    case V_UPLOADED = 'uploaded';
    case V_FAILED = 'failed';
}
