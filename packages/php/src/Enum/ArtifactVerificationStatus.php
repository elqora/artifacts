<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Enum;

enum ArtifactVerificationStatus: string
{
    case V_PENDING = 'pending';
    case V_VERIFIED = 'verified';
    case V_REJECTED = 'rejected';
    case V_WAIVED = 'waived';
}
