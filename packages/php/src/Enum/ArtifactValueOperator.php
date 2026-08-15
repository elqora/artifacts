<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\Artifact\Enum;

enum ArtifactValueOperator: string
{
    case V_EQ = 'eq';
    case V_NEQ = 'neq';
    case V_GT = 'gt';
    case V_GTE = 'gte';
    case V_LT = 'lt';
    case V_LTE = 'lte';
    case V_CONTAINS = 'contains';
    case V_IN = 'in';
}
