import type { Handle } from '@sveltejs/kit';
import { sequence } from '@sveltejs/kit/hooks';
import { i18n } from '$lib/i18n';
import { withClerkHandler } from 'svelte-clerk/server';

const handleParaglide: Handle = i18n.handle();
export const handle: Handle = sequence(handleParaglide, withClerkHandler())
