import { useState } from "react";
import { CreateDiscount, CreateDiscountSchema } from "@/types/discount";
import { useMutationCreateDiscount } from "@/query/discount";

interface SeasonalFormProps {
  selectedDiscountCategoryId: number;
}

const SeasonalForm = ({ selectedDiscountCategoryId }: SeasonalFormProps) => {
  const [everyXTHB, setEveryXTHB] = useState<number>(0);
  const [discountYTHB, setDiscountYTHB] = useState<number>(0);
  const { mutateAsync, isPending, error } = useMutationCreateDiscount();

  const onSubmitDiscount = async (e: React.FormEvent) => {
    e.preventDefault();

    const discountData: CreateDiscount = {
      seasonal: {
        every_xthb : everyXTHB,
        discount_ythb :discountYTHB,
      },
      discount_category_id: selectedDiscountCategoryId,
    };

    try {
      await mutateAsync(discountData); 
    } catch (err) {
      console.error("Error creating seasonal discount", err);
    }
  };

  return (
    <form onSubmit={onSubmitDiscount}>
      <h2 className="text-xl font-bold text-gray-800 mb-4">Seasonal Discount</h2>

      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-2">
          Every X THB
        </label>
        <input
          type="number"
          name="everyXTHB"
          value={everyXTHB}
          onChange={(e) => setEveryXTHB(Number(e.target.value))}
          placeholder="Enter amount (X)"
          className="block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
          required
        />
      </div>

      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-2">
          Discount Y THB
        </label>
        <input
          type="number"
          name="discountYTHB"
          value={discountYTHB}
          onChange={(e) => setDiscountYTHB(Number(e.target.value))}
          placeholder="Enter discount amount (Y)"
          className="block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
          required
        />
      </div>

      {error && <p className="text-red-500 text-sm mb-4">{error.message}</p>} {/* Display error message */}

      <button
        type="submit"
        className="w-full bg-blue-500 hover:bg-blue-600 text-white font-medium py-2 px-4 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        disabled={isPending}
      >
        {isPending ? "Creating..." : "Apply Seasonal Discount"}
      </button>
    </form>
  );
};

export default SeasonalForm;
