package data

type amountDef struct {
	Label     string
	Amount    int64
	Frequency Frequency
}

//Account holds balance informations
type Account struct {
	AccountID            string
	AvailableAmount      int64
	Revenues             []*Revenue
	Expenses             []*Expense
	Savings              int64
	totalMonthlyExpenses int64
	totalMonthlyRevenue  int64
}

//ExpenseCat expense category.
type ExpenseCat string

const (
	//Food expenses
	Food ExpenseCat = "Food"
	//Clothes expenses
	Clothes = "Clothes"
	//FastFood and restaurants
	FastFood = "FastFood"
	//Leisure and hobbies
	Leisure = "Leasure"
)

//ExpenseType expense type.
type ExpenseType string

const (
	//Irreducible expenses such as rent or electric bill
	Irreducible ExpenseType = "Irreducible"
	//Reducible expenses such as insurance, subscriptions ( cable, VOD ...)
	Reducible = "Reducible"
)

//Frequency type for amount definition
type Frequency string

const (
	//Monthly once a month ( revenue like salary or expense like rent.)
	Monthly Frequency = "Monthly"
	//Annually once a year.
	Annually = "Annualy"
	//Once one shot.
	Once = "Once"
)

func (amountDef *amountDef) equals(amountDef2 amountDef) bool {
	if amountDef.Amount != amountDef2.Amount || amountDef.Label != amountDef2.Label {
		return false
	}
	return true
}

//Revenue type
type Revenue struct {
	AmountDef amountDef
}

//Expense type
type Expense struct {
	amountDef       amountDef
	expenseCategory ExpenseCat
	expenseType     ExpenseType
}

func (expense *Expense) equals(expense2 *Expense) bool {
	amountDef := expense.amountDef
	amountDef2 := expense2.amountDef
	if !(amountDef.equals(amountDef2) || expense.expenseType != expense2.expenseType || expense.expenseCategory != expense2.expenseCategory) {
		return false
	}
	return true
}

func (account *Account) getTotalRevenue(frequency Frequency) int64 {
	var sum int64 = 0
	for _, v := range account.Revenues {
		if v.AmountDef.Frequency == frequency {
			sum += v.AmountDef.Amount
		}
	}
	return sum
}

func (account *Account) getTotalExpenses(frequency Frequency) int64 {
	var sum int64 = 0
	for _, v := range account.Expenses {
		if v.amountDef.Frequency == frequency {
			sum += v.amountDef.Amount
		}
	}
	return sum
}

//RemoveExpense from the user account
func (account *Account) RemoveExpense(expense Expense) {
	index := 0
	for idx, exp := range account.Expenses {
		if expense.equals(exp) {
			index = idx
			break
		}
	}
	if index != 0 {
		account.Expenses = append(account.Expenses[:index], account.Expenses[index+1:]...)
	}
}

//RemoveRevenue from the user account
func (account *Account) RemoveRevenue(revenue Revenue) {
	index := 0
	for idx, rev := range account.Revenues {
		if rev.AmountDef.equals(revenue.AmountDef) {
			index = idx
		}
	}
	if index != 0 {
		account.Revenues = append(account.Revenues[:index], account.Revenues[index+1:]...)
	}
}
