import { z } from "zod";

export const FixedAmountDiscountSchema = z.object({
  id: z.number().int().optional(),
  amount: z.number(),
});

export const CreateFixedAmountDiscountSchema = z.object({
  amount: z.number(),
});

export type FixedAmountDiscount = z.infer<typeof FixedAmountDiscountSchema>;
export type CreateFixedAmountDiscount = z.infer<typeof CreateFixedAmountDiscountSchema>;