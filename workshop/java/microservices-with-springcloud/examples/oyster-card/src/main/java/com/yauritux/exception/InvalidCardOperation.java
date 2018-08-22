package com.yauritux.exception;

import com.yauritux.model.constant.CardExceptionType;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
public class InvalidCardOperation extends CardException {

	private static final long serialVersionUID = -1744103431552901326L;
	
	private CardExceptionType exceptionType;
	
	public InvalidCardOperation(String message) {
		super(message);
	}
	
	public InvalidCardOperation(String message, CardExceptionType exceptionType) {
		super(message);
		this.exceptionType = exceptionType;
	}
	
	public void setExceptionType(CardExceptionType exceptionType) {
		this.exceptionType = exceptionType;
	}
	
	public CardExceptionType getExceptionType() {
		return exceptionType;
	}
}
