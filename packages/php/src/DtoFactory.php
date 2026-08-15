<?php

declare(strict_types=1);

namespace Elqora\Artifact;

use Elqora\Artifact\Dto\ArtifactActorCondition;
use Elqora\Artifact\Dto\ArtifactAndCondition;
use Elqora\Artifact\Dto\ArtifactExistsCondition;
use Elqora\Artifact\Dto\ArtifactNotCondition;
use Elqora\Artifact\Dto\ArtifactOrCondition;
use Elqora\Artifact\Dto\ArtifactStateCondition;
use Elqora\Artifact\Dto\ArtifactValueCondition;
use Elqora\Artifact\Dto\AudioArtifactValueSchema;
use Elqora\Artifact\Dto\BooleanArtifactValueSchema;
use Elqora\Artifact\Dto\CollectionArtifactValueSchema;
use Elqora\Artifact\Dto\CurrencyArtifactValueSchema;
use Elqora\Artifact\Dto\DateArtifactValueSchema;
use Elqora\Artifact\Dto\DatetimeArtifactValueSchema;
use Elqora\Artifact\Dto\FileArtifactValueSchema;
use Elqora\Artifact\Dto\HostedArtifactSource;
use Elqora\Artifact\Dto\ImageArtifactValueSchema;
use Elqora\Artifact\Dto\InlineArtifactSource;
use Elqora\Artifact\Dto\LinkArtifactValueSchema;
use Elqora\Artifact\Dto\LocalArtifactSource;
use Elqora\Artifact\Dto\LocationArtifactValueSchema;
use Elqora\Artifact\Dto\NumberArtifactValueSchema;
use Elqora\Artifact\Dto\ObjectArtifactSource;
use Elqora\Artifact\Dto\ProviderArtifactSource;
use Elqora\Artifact\Dto\ReferenceArtifactValueSchema;
use Elqora\Artifact\Dto\SignatureArtifactValueSchema;
use Elqora\Artifact\Dto\StructuredArtifactValueSchema;
use Elqora\Artifact\Dto\TextArtifactValueSchema;
use Elqora\Artifact\Dto\TimeArtifactValueSchema;
use Elqora\Artifact\Dto\UrlArtifactSource;
use Elqora\Artifact\Dto\VideoArtifactValueSchema;
use InvalidArgumentException;

final class DtoFactory
{
    /** @param array<string, mixed> $data */
    public static function source(array $data): object
    {
        return match ($data['type'] ?? null) {
            'inline' => InlineArtifactSource::fromArray($data),
            'local' => LocalArtifactSource::fromArray($data),
            'object' => ObjectArtifactSource::fromArray($data),
            'url' => UrlArtifactSource::fromArray($data),
            'hosted' => HostedArtifactSource::fromArray($data),
            'provider' => ProviderArtifactSource::fromArray($data),
            default => throw new InvalidArgumentException('Unsupported ArtifactSource discriminator.'),
        };
    }

    /** @param array<string, mixed> $data */
    public static function condition(array $data): object
    {
        return match ($data['kind'] ?? null) {
            'state' => ArtifactStateCondition::fromArray($data),
            'actor' => ArtifactActorCondition::fromArray($data),
            'artifact_exists' => ArtifactExistsCondition::fromArray($data),
            'artifact_value' => ArtifactValueCondition::fromArray($data),
            'and' => ArtifactAndCondition::fromArray($data),
            'or' => ArtifactOrCondition::fromArray($data),
            'not' => ArtifactNotCondition::fromArray($data),
            default => throw new InvalidArgumentException('Unsupported ArtifactCondition discriminator.'),
        };
    }

    /** @param array<string, mixed> $data */
    public static function valueSchema(array $data): object
    {
        return match ($data['valueType'] ?? null) {
            'text' => TextArtifactValueSchema::fromArray($data),
            'number' => NumberArtifactValueSchema::fromArray($data),
            'boolean' => BooleanArtifactValueSchema::fromArray($data),
            'currency' => CurrencyArtifactValueSchema::fromArray($data),
            'date' => DateArtifactValueSchema::fromArray($data),
            'datetime' => DatetimeArtifactValueSchema::fromArray($data),
            'time' => TimeArtifactValueSchema::fromArray($data),
            'location' => LocationArtifactValueSchema::fromArray($data),
            'file' => FileArtifactValueSchema::fromArray($data),
            'image' => ImageArtifactValueSchema::fromArray($data),
            'video' => VideoArtifactValueSchema::fromArray($data),
            'audio' => AudioArtifactValueSchema::fromArray($data),
            'link' => LinkArtifactValueSchema::fromArray($data),
            'structured' => StructuredArtifactValueSchema::fromArray($data),
            'reference' => ReferenceArtifactValueSchema::fromArray($data),
            'signature' => SignatureArtifactValueSchema::fromArray($data),
            'collection' => CollectionArtifactValueSchema::fromArray($data),
            default => throw new InvalidArgumentException('Unsupported ArtifactValueSchema discriminator.'),
        };
    }
}
