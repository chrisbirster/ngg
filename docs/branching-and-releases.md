# Branching and releases

1. Branch `feature/*` from `dev`.
2. Open a pull request into `dev`.
3. Require CI before merge.
4. Prepare releases through a pull request from `dev` to `main`.
5. Tag the merge commit with semantic versioning.
6. Deploy only release tags to production.

Hotfixes branch from `main`, merge back to both `main` and `dev`.
