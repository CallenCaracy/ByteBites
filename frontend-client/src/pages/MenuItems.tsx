import { useNavigate, useParams } from "react-router-dom";
import { useQuery } from "@apollo/client";
import { GET_MENU_ITEM_BY_ID } from "../graphql/Menuqueries";
import Navbar from "../components/NavBar";
import placeholder from "../assets/placeholder.jpg";

const MenuItem: React.FC = () => {
    const { menuId } = useParams();
    const navigate = useNavigate();

    const { data, loading, error } = useQuery(GET_MENU_ITEM_BY_ID, {
        variables: { id: menuId },
        skip: !menuId,
    });

    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error fetching menu data: {error.message}</p>;
    if (!data || !data.getMenuItemById) return <p>No menu data found.</p>;

    const item = data.getMenuItemById;

    return (
        <>
            <Navbar />
            <div className="p-8 max-w-xl mx-auto bg-white rounded-2xl shadow-[0_4px_10px_0_rgba(0,0,0,0.5)] space-y-5">
                <img src={placeholder} alt={item.name} className="w-full rounded-xl object-cover h-64" />

                <div className="flex justify-between items-center">
                    <h2 className="text-gray-900 text-2xl font-bold">{item.name}</h2>
                    <p className="text-xl font-semibold text-green-700 whitespace-nowrap">
                        ₱{item.price.toFixed(2)}
                    </p>
                </div>

                <p className="text-gray-600">{item.description}</p>
                <p className="text-sm text-gray-500">Category: {item.category}</p>
                <p className="text-sm text-gray-500">
                    Status:{" "}
                    <span className={item.availability_status ? "text-green-600 font-semibold" : "text-red-600 font-semibold"}>
                        {item.availability_status ? "Available" : "Out of stock"}
                    </span>
                </p>

                <div className="flex gap-4 pt-4">
                    <button
                        className="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition"
                    >
                        Add to Cart
                    </button>
                    <button
                        className="px-4 py-2 bg-gray-300 text-white rounded-lg hover:bg-gray-400 transition"
                        onClick={() => navigate("/dashboard")}
                    >
                        Back
                    </button>
                </div>
            </div>
        </>
    );
};

export default MenuItem;