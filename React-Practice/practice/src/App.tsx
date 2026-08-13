//======without props============


// type Employee = {
//   firstName: string;
//   lastName: string;
//   employeeId: string;
//   designation: string;
// };

// function Header() {
//   return (
//     <header>
//       <h1>Users</h1>
//       <button>Register</button>
//     </header>
//   );
// }

// function UserCard() {
//   const employee: Employee = {
//     firstName: "Chandini",
//     lastName: "Oleti",
//     employeeId: "EMP101",
//     designation: "Software Engineer"
//   };

//   return (
//     <div>
//       <h2>
//         {employee.firstName} {employee.lastName}
//       </h2>

//       <p>Employee ID: {employee.employeeId}</p>

//       <p>Designation: {employee.designation}</p>
//     </div>
//   );
// }

// function UserList() {
//   return (
//     <main>
//       <UserCard />
//     </main>
//   );
// }

// function App() {
//   return (
//     <div>
//       <Header />
//       <UserList />
//     </div>
//   );
// }

// export default App;

//=========with props=============


// type Employee = {
//   firstName: string;
//   lastName: string;
//   employeeId: string;
//   designation: string;
// };

// type UserCardProps = {
//   employee: Employee;
// };

// function UserCard({ employee }: UserCardProps) {
//   return (
//     <div>
//       <h2>
//         {employee.firstName} {employee.lastName}
//       </h2>

//       <p>Employee ID: {employee.employeeId}</p>

//       <p>Designation: {employee.designation}</p>
//     </div>
//   );
// }

// function App() {
//   const employee: Employee = {
//     firstName: "Chandini",
//     lastName: "Oleti",
//     employeeId: "EMP101",
//     designation: "Software Engineer"
//   };

//   return (
//     <div>
//       <h1>Users</h1>

//       <UserCard employee={employee} />
//     </div>
//   );
// }

// export default App;


//===============useState===================

// import { useState } from "react";

// function App() {
//   const [count, setCount] = useState(0);

//   return (
//     <div>
//       <h1>{count}</h1>

//       <button onClick={() => setCount(count + 1)}>
//         Increase
//       </button>
//     </div>
//   );
// }

// export default App;


//==========small registration form============


// import { useState } from "react";

// type Employee = {
//   firstName: string;
//   lastName: string;
//   employeeId: string;
//   designation: string;
// };

// function App() {
//   const [employees, setEmployees] = useState<Employee[]>([]);

//   function addEmployee() {
//     const newEmployee: Employee = {
//       firstName: "Chandini",
//       lastName: "Oleti",
//       employeeId: "EMP101",
//       designation: "Software Engineer"
//     };

//     setEmployees([...employees, newEmployee]);
//   }

//   return (
//     <div>
//       <h1>Users</h1>

//       {employees.length === 0 ? (
//         <p>No users present</p>
//       ) : (
//         employees.map((employee) => (
//           <div key={employee.employeeId}>
//             <h2>
//               {employee.firstName} {employee.lastName}
//             </h2>

//             <p>Employee ID: {employee.employeeId}</p>

//             <p>Designation: {employee.designation}</p>
//           </div>
//         ))
//       )}

//       <button onClick={addEmployee}>
//         Add User
//       </button>
//     </div>
//   );
// }

// export default App;


//==========event on click==========


// function App() {
//   function handleClick() {
//     console.log("Register clicked");
//   }

//   return (
//     <div>
//       <h1>Users</h1>

//       <button onClick={handleClick}>
//         Register
//       </button>
//     </div>
//   );
// }

// export default App;


//==================event +state ==============

// import { useState } from "react";

// function App() {
//   const [message, setMessage] = useState("No users present");

//   function handleRegister() {
//     setMessage("Registration form opened");
//   }

//   return (
//     <div>
//       <h1>Users</h1>

//       <p>{message}</p>

//       <button onClick={handleRegister}>
//         Register
//       </button>
//     </div>
//   );
// }

// export default App;

//=====state + input


// import { useState } from "react";

// function App() {
//   const [firstName, setFirstName] = useState("");

//   function handleFirstNameChange(
//     event: React.ChangeEvent<HTMLInputElement>
//   ) {
//     setFirstName(event.target.value);
//   }

//   function handleSubmit(
//     event: React.FormEvent<HTMLFormElement>
//   ) {
//     event.preventDefault();

//     console.log("Registered:", firstName);
//   }

//   return (
//     <div>
//       <h1>Register Employee</h1>

//       <form onSubmit={handleSubmit}>
//         <div>
//           <label>First Name</label>

//           <input
//             value={firstName}
//             onChange={handleFirstNameChange}
//           />
//         </div>

//         <button type="submit">
//           Register
//         </button>
//       </form>
//     </div>
//   );
// }

// export default App;


//============small page=============



// import { useState } from "react";

// type RegistrationForm = {
//   firstName: string;
//   lastName: string;
//   email: string;
//   gender: string;
//   designation: string;
//   experience: number;
// };

// function App() {
//   const [form, setForm] = useState<RegistrationForm>({
//     firstName: "",
//     lastName: "",
//     email: "",
//     gender: "",
//     designation: "",
//     experience: 0
//   });

//   function handleChange(
//     event: React.ChangeEvent<
//       HTMLInputElement | HTMLSelectElement
//     >
//   ) {
//     const { name, value } = event.target;

//     setForm({
//       ...form,
//       [name]:
//         name === "experience"
//           ? Number(value)
//           : value
//     });
//   }

//   return (
//     <div>
//       <h1>Employee Registration</h1>

//       <div>
//         <label htmlFor="firstName">
//           First Name
//         </label>

//         <input
//           id="firstName"
//           name="firstName"
//           type="text"
//           value={form.firstName}
//           onChange={handleChange}
//         />
//       </div>

//       <div>
//         <label htmlFor="lastName">
//           Last Name
//         </label>

//         <input
//           id="lastName"
//           name="lastName"
//           type="text"
//           value={form.lastName}
//           onChange={handleChange}
//         />
//       </div>

//       <div>
//         <label htmlFor="email">
//           Email
//         </label>

//         <input
//           id="email"
//           name="email"
//           type="email"
//           value={form.email}
//           onChange={handleChange}
//         />
//       </div>

//       <div>
//         <label htmlFor="gender">
//           Gender
//         </label>

//         <select
//           id="gender"
//           name="gender"
//           value={form.gender}
//           onChange={handleChange}
//         >
//           <option value="">
//             Select gender
//           </option>

//           <option value="male">
//             Male
//           </option>

//           <option value="female">
//             Female
//           </option>

//           <option value="other">
//             Other
//           </option>
//         </select>
//       </div>

//       <div>
//         <label htmlFor="designation">
//           Designation
//         </label>

//         <select
//           id="designation"
//           name="designation"
//           value={form.designation}
//           onChange={handleChange}
//         >
//           <option value="">
//             Select designation
//           </option>

//           <option value="software-engineer">
//             Software Engineer
//           </option>

//           <option value="developer">
//             Developer
//           </option>

//           <option value="manager">
//             Manager
//           </option>
//         </select>
//       </div>

//       <div>
//         <label htmlFor="experience">
//           Experience
//         </label>

//         <input
//           id="experience"
//           name="experience"
//           type="number"
//           min="0"
//           value={form.experience}
//           onChange={handleChange}
//         />
//       </div>

//       <h2>Preview</h2>

//       <p>
//         Name: {form.firstName} {form.lastName}
//       </p>

//       <p>Email: {form.email}</p>

//       <p>Gender: {form.gender}</p>

//       <p>Designation: {form.designation}</p>

//       <p>Experience: {form.experience}</p>
//     </div>
//   );
// }

// export default App;


//==============form with validations===========


import { useState } from "react";

type RegistrationForm = {
  firstName: string;
  lastName: string;
  email: string;
  phone: string;
  employeeId: string;
  experience: number;
  designation: string;
};

type FormErrors = Partial<
  Record<keyof RegistrationForm, string>
>;

function App() {
  const [form, setForm] = useState<RegistrationForm>({
    firstName: "",
    lastName: "",
    email: "",
    phone: "",
    employeeId: "",
    experience: 0,
    designation: ""
  });

  const [errors, setErrors] = useState<FormErrors>({});

  function handleChange(
    event: React.ChangeEvent<
      HTMLInputElement | HTMLSelectElement
    >
  ) {
    const { name, value } = event.target;

    setForm({
      ...form,
      [name]:
        name === "experience"
          ? Number(value)
          : value
    });
  }

  function validateForm(): FormErrors {
    const errors: FormErrors = {};

    if (!form.firstName.trim()) {
      errors.firstName =
        "First Name is required";
    }

    if (!form.lastName.trim()) {
      errors.lastName =
        "Last Name is required";
    }

    const emailPattern =
      /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

    if (!form.email.trim()) {
      errors.email = "Email is required";
    } else if (!emailPattern.test(form.email)) {
      errors.email =
        "Enter a valid email address";
    }

    const phonePattern = /^[6-9]\d{9}$/;

    if (!phonePattern.test(form.phone)) {
      errors.phone =
        "Enter a valid 10-digit phone number";
    }

    const employeeIdPattern = /^EMP\d{3}$/;

    if (!employeeIdPattern.test(form.employeeId)) {
      errors.employeeId =
        "Employee ID must be like EMP101";
    }

    if (
      form.experience < 0 ||
      form.experience > 50
    ) {
      errors.experience =
        "Experience must be between 0 and 50";
    }

    if (!form.designation) {
      errors.designation =
        "Please select a designation";
    }

    return errors;
  }

  function handleSubmit(
    event: React.FormEvent<HTMLFormElement>
  ) {
    event.preventDefault();

    const validationErrors = validateForm();

    if (Object.keys(validationErrors).length > 0) {
      setErrors(validationErrors);
      return;
    }

    setErrors({});

    console.log("Registration successful");
    console.log(form);
  }

  return (
    <div>
      <h1>Employee Registration</h1>

      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="firstName">
            First Name
          </label>

          <input
            id="firstName"
            name="firstName"
            value={form.firstName}
            onChange={handleChange}
          />

          {errors.firstName && (
            <p>{errors.firstName}</p>
          )}
        </div>

        <div>
          <label htmlFor="lastName">
            Last Name
          </label>

          <input
            id="lastName"
            name="lastName"
            value={form.lastName}
            onChange={handleChange}
          />

          {errors.lastName && (
            <p>{errors.lastName}</p>
          )}
        </div>

        <div>
          <label htmlFor="email">
            Email
          </label>

          <input
            id="email"
            name="email"
            type="email"
            value={form.email}
            onChange={handleChange}
          />

          {errors.email && (
            <p>{errors.email}</p>
          )}
        </div>

        <div>
          <label htmlFor="phone">
            Phone
          </label>

          <input
            id="phone"
            name="phone"
            type="tel"
            value={form.phone}
            onChange={handleChange}
          />

          {errors.phone && (
            <p>{errors.phone}</p>
          )}
        </div>

        <div>
          <label htmlFor="employeeId">
            Employee ID
          </label>

          <input
            id="employeeId"
            name="employeeId"
            value={form.employeeId}
            onChange={handleChange}
          />

          {errors.employeeId && (
            <p>{errors.employeeId}</p>
          )}
        </div>

        <div>
          <label htmlFor="experience">
            Experience
          </label>

          <input
            id="experience"
            name="experience"
            type="number"
            min="0"
            value={form.experience}
            onChange={handleChange}
          />

          {errors.experience && (
            <p>{errors.experience}</p>
          )}
        </div>

        <div>
          <label htmlFor="designation">
            Designation
          </label>

          <select
            id="designation"
            name="designation"
            value={form.designation}
            onChange={handleChange}
          >
            <option value="">
              Select designation
            </option>

            <option value="software-engineer">
              Software Engineer
            </option>

            <option value="developer">
              Developer
            </option>

            <option value="manager">
              Manager
            </option>
          </select>

          {errors.designation && (
            <p>{errors.designation}</p>
          )}
        </div>

        <button type="submit">
          Register
        </button>
      </form>
    </div>
  );
}

export default App;
