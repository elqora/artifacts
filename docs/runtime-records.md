# Submissions and verification

Submissions preserve operational history instead of destructively replacing a
value. They identify the artifact, may pin the exact artifact version, identify
the submitting actor, and can carry an inline value plus time, location, device,
network, user-agent, and host metadata context.

Verification is a first-class record with closed status and open method. It may
reference an artifact, exact version, and submission. Approval or rejection of
one version never applies automatically to later versions. Hosts should pin
evidence-sensitive decisions and verify cross-record identifier consistency.

Submission-context fields can be sensitive or regulated. The schema describes
their shape; hosts decide whether collection is safe, lawful, and necessary.
