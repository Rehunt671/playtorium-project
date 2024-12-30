import { z } from "zod";

export const SeasonalDiscountSchema = z.object({
  id: z.number().int(),
  every_xthb: z.number(),
  discount_ythb: z.number(),
});

export const CreateSeasonalDiscountSchema = z.object({
  every_xthb: z.number(),
  discount_ythb: z.number(),
});

export type SeasonalDiscount = z.infer<typeof SeasonalDiscountSchema>;
export type CreateSeasonalDiscount = z.infer<typeof CreateSeasonalDiscountSchema>;
