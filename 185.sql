select Department.Name Department, e1.Name Employee, e1.Salary Salary
from Employee e1,
     Department
where (
          select count(distinct e2.Salary)
          from Employee e2
          where e1.DepartmentId = e2.DepartmentId
            and e2.Salary > e1.Salary
      ) < 3
  and e1.DepartmentId = Department.Id
order by Department.Name, e1.Salary desc