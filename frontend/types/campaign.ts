import { z } from 'zod';

export const CampaignSchema = z.object({
  id: z.number(),
  name: z.string(),
});

export type Campaign = z.infer<typeof CampaignSchema>;
