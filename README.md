# NavKode 🌍  
An interactive web and mobile app to improve navigation and accessibility in Kozhikode by providing real-time route suggestions that factor in traffic congestion, parking availability, and nearby dining options.  

---

## 🚀 Features  
- **Landmark-Based Navigation**: Select from major landmarks in Kozhikode and get the best route.  
- **Real-Time Traffic**: Avoid congested routes with live traffic data.  
- **Parking Information**: Locate nearby parking spots with availability details.  
- **Restaurant Recommendations**: Discover popular food options near your destination.  
- **Google Maps Integration**: Redirect to Google Maps for turn-by-turn navigation.  

---

## 🛠️ Tech Stack  

### **Frontend**  
- HTML, CSS, Tailwind CSS  
- Svelte for reactivity  
- TypeScript for logic  

### **Backend**  
- Go Lang for REST API  
- PostgreSQL for database  
- Traefik for load balancing   

### **Other Integrations**   
- OpenStreetMap + MapLibre for custom mapping  
- Cloudflare for security  
- GitHub for CI/CD  

---

## 📋 Installation  

### Prerequisites  
- Node.js for frontend developmen
- Go 1.22+

### Steps  
1. Clone the repository:  
   ```bash
   git clone https://github.com/hxri-nxrxyxn/navkode/
   cd navkode
   ```

2. Set up the backend:  
   ```bash
   cd api
   go run main.go
   ```

3. Set up the frontend:  
   ```bash
   npm install
   npm run dev
   ```

4. Access the app:  
   - **Web**: `http://localhost:5173`  
   - **Mobile**: Convert the app to APK using CapacitorJS.  

---

## 🌟 How It Works  

1. **Select a Landmark**: Choose a destination from the predefined list.  
2. **Analyze Routes**: The app calculates the best route based on traffic, parking, and restaurant data.  
3. **Navigate**: Confirm the route and get redirected to Google Maps for step-by-step guidance.  

---

## 📅 Upcoming Features  
- **Crowdsourced Traffic Reporting**: Allow users to report traffic conditions in real-time.  
- **AI-Powered Suggestions**: Use machine learning to improve route recommendations.  
- **Offline Navigation**: Download maps for use without internet connectivity.  

---

## **API Documentation**

### **Base URL:**  
`http://<your-domain>/api/v1`

---

### **Authentication**

#### **1. Create User**  
**POST** `/user`  

- **Description:** Registers a new user.  
- **Request Body:**  
  ```json
  {
    "email": "string (required)",
    "password": "string (required)"
  }
  ```
- **Responses:**  
  - **201:** User created successfully.  
  - **400:** Error parsing request body.  
  - **422:** Email and password are required.  
  - **409:** Email already taken.  
  - **501:** Error inserting into the database.  

---

#### **2. Login**  
**POST** `/login`  

- **Description:** Authenticates a user with email and password.  
- **Request Body:**  
  ```json
  {
    "email": "string (required)",
    "password": "string (required)"
  }
  ```
- **Responses:**  
  - **200:** Successfully logged in.  
  - **400:** Error parsing request body.  
  - **422:** Email and password are required.  
  - **404:** User not found.  
  - **500:** Password is incorrect or error retrieving user.  

---

#### **3. Get User by ID**  
**GET** `/user/:id`  

- **Description:** Retrieves user details by ID.  
- **Path Parameters:**  
  - `id` (required): User ID.  
- **Responses:**  
  - **200:** Successfully retrieved user.  
  - **404:** User not found.  
  - **500:** Error retrieving user.  

---

### **Routes**

#### **4. Create Routes**  
**POST** `/routes`  

- **Description:** Adds a new route.  
- **Request Body:**  
  ```json
  {
    "field1": "value",
    "field2": "value"
    }
  ```
- **Responses:**  
  - **201:** Routes created successfully.  
  - **400:** Error parsing request body.  
  - **501:** Error inserting into the database.  

#### **5. Update Routes**  
**PATCH** `/routes/:id`  

- **Description:** Updates route information.  
- **Path Parameters:**  
  - `id` (required): Route ID.  
- **Request Body:**  
  ```json
  {
    "field_to_update": "new_value"
  }
  ```
- **Responses:**  
  - **200:** Routes updated successfully.  
  - **400:** Error parsing request body.  
  - **500:** Error finding or updating the route.  

---

### **Restaurants**

#### **6. Create Restaurant**  
**POST** `/restaurant`  

- **Description:** Adds a new restaurant.  
- **Request Body:**  
  ```json
  {
    "name": "string (required)",
    "location": "string (required)",
    "category": "string"
  }
  ```
- **Responses:**  
  - **201:** Restaurant created successfully.  
  - **400:** Error parsing request body.  
  - **501:** Error inserting into the database.  

---

### **Parking**

#### **7. Create Parking**  
**POST** `/parking`  

- **Description:** Adds new parking details.  
- **Request Body:**  
  ```json
  {
    "name": "string (required)",
    "location": "string (required)",
    "capacity": "integer (required)"
  }
  ```
- **Responses:**  
  - **201:** Parking created successfully.  
  - **400:** Error parsing request body.  
  - **501:** Error inserting into the database.  

---

### **Landmarks**

#### **8. Create Landmark**  
**POST** `/landmark`  

- **Description:** Adds a new landmark.  
- **Request Body:**  
  ```json
  {
    "name": "string (required)",
    "location": "string (required)",
    "category": "string (optional)"
  }
  ```
- **Responses:**  
  - **201:** Landmark created successfully.  
  - **400:** Error parsing request body.  
  - **501:** Error inserting into the database.  

#### **9. Get Landmark by ID**  
**GET** `/landmark/:id`  

- **Description:** Retrieves landmark details by ID.  
- **Path Parameters:**  
  - `id` (required): Landmark ID.  
- **Responses:**  
  - **200:** Successfully retrieved landmark.  
  - **401:** Error retrieving the landmark.  

#### **10. Get Landmarks by Category**  
**GET** `/landmarkCat/:category`  

- **Description:** Retrieves landmarks by category.  
- **Path Parameters:**  
  - `category` (required): Landmark category.  
- **Responses:**  
  - **200:** Successfully retrieved landmarks.  
  - **401:** Error retrieving landmarks.  

---

### **Error Responses**  

- **400:** Bad Request (Invalid or malformed request data).  
- **401:** Unauthorized access.  
- **404:** Resource not found.  
- **422:** Validation errors.  
- **500:** Internal Server Error.  
- **501:** Database insertion error.  

--- 

## 📧 Contact

For any questions or inquiries, please contact [me](hari@laddu.cc). 📬
