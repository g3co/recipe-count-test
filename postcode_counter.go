package main

type PostCodeCounter struct {
	PostCode postCodeName
	Start    int64
	End      int64
	Counter  int64
}

// NewPostCodeCounter returns new instance of PostCodeCounter obj
func NewPostCodeCounter(pc string, start, end int64) PostCodeCounter {
	postCode := postCodeName{}
	copy(postCode[:], pc)

	return PostCodeCounter{
		PostCode: postCode,
		Start:    start,
		End:      end,
		Counter:  0,
	}
}

// Check checks that postcode and time is valid and increments counter
func (c *PostCodeCounter) Check(pc postCodeName, start, end int64) {
	if pc == c.PostCode {
		if start <= c.Start && end >= c.End {
			c.Counter++
			return
		}

		if start > end {
			if start <= c.Start+24 && end >= c.End {
				c.Counter++
				return
			}

			if start <= c.Start && end+24 >= c.End {
				c.Counter++
				return
			}
		}
	}
}

// GetStrPostCode returns string representation of the PostCode
func (c *PostCodeCounter) GetStrPostCode() string {
	return fixedSizeArrToString(c.PostCode[:])
}
