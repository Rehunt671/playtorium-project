import {z} from "zod";

export const PercentageDiscountSchema = z.object({
  id: z.number().int(),
  percentage: z.number(),
});


export const CreatePercentageDiscountSchema = z.object({
    percentage: z.number(),
});
  
export type PercentageDiscount = z.infer<typeof PercentageDiscountSchema>;
export type CreatePercentageDiscount = z.infer<typeof CreatePercentageDiscountSchema>;