import { z } from "zod";
export const DiscountCategorySchema = z.object({
  id: z.number().int(),
  name: z.string(),
});

export type DiscountCategory = z.infer<typeof DiscountCategorySchema>;
