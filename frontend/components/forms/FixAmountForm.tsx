import { useState } from "react";
import { CreateDiscount} from "@/types/discount";
import { useMutationCreateDiscount } from "@/query/discount";

interface FixedAmountFormProps {
  selectedDiscountCategoryId: number; 
}

const FixedAmountForm = ({ selectedDiscountCategoryId }: FixedAmountFormProps) => {
  const [amount, setAmount] = useState<number>(0);
  const { mutateAsync, isPending } = useMutationCreateDiscount();
  const [error, setError] = useState('');

  const onSubmitDiscount = async (e: React.FormEvent) => {
    e.preventDefault();

    if (amount === 0) {
      setError("Discount amount must > 0");
      return; 
    }

    const discountData: CreateDiscount = {
      fixed_amount: {
        amount: amount,
      },
      discount_category_id: selectedDiscountCategoryId,
    };

    try {
      await mutateAsync(discountData); 
      setError('')
    } catch (err) {
      console.error("Error creating discount", err);
    }
  };

  return (
    <form onSubmit={onSubmitDiscount}>
      <h2 className="text-xl font-bold text-gray-800 mb-4">Fixed Amount Discount</h2>

      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-2">
          Discount Amount
        </label>
        <input
          type="number"
          name="amount"
          value={amount}
          onChange={(e) => setAmount(Number(e.target.value))}
          placeholder="Enter the discount amount"
          className="block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
          required
        />
      </div>

      {error && <p className="text-red-500 text-sm mb-4">{error}</p>} {/* Display error message */}

      <button
        type="submit"
        className="w-full bg-blue-500 hover:bg-blue-600 text-white font-medium py-2 px-4 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        disabled={isPending}
      >
        {isPending ? "Creating..." : "Apply Discount"}
      </button>
    </form>
  );
};

export default FixedAmountForm;
