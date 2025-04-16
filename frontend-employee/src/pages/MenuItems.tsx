import { useParams } from "react-router-dom";
import { useQuery } from "@apollo/client";
import { GET_MENU_ITEM_BY_ID } from "../graphql/Menuqueries";
import Navbar from "../components/NavBar";

const MenuItem: React.FC = () => {
    const { menuId } = useParams();

    const { data, loading, error } = useQuery(GET_MENU_ITEM_BY_ID, {
        variables: { id: menuId },
        skip: !menuId,
    });

    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error fetching menu data: {error.message}</p>;
    if (!data || !data.getMenuItemById) return <p>No menu data found.</p>;

    const item = data.getMenuItemById;

    return (
        <><Navbar /><div className="p-6 max-w-md mx-auto bg-white rounded-xl shadow-md space-y-4">
            <button>
                Back
            </button>
            <img src={item.image_url} alt={item.name} className="w-full rounded" />
            <h2 className="text-gray-900 text-2xl font-bold">{item.name}</h2>
            <p className="text-gray-600">{item.description}</p>
            <p className="text-xl font-semibold text-green-700">${item.price.toFixed(2)}</p>
            <p className="text-sm text-gray-500">Category: {item.category}</p>
            <p className="text-sm text-gray-500">
                Status: {item.availability_status ? "Available" : "Out of stock"}
            </p>
            <button>
                Update
            </button>
            <button>
                Delete
            </button>
        </div></>
    );
};
  
  export default MenuItem;  