package com.yauritux.exception;

import com.yauritux.model.constant.CardExceptionType;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
public class CardNotFoundException extends CardException {

	private static final long serialVersionUID = -6079926321282686121L;
	
	private final CardExceptionType exceptionType = CardExceptionType.CARD_NOT_FOUND_EXCEPTION;

	public CardNotFoundException(String message) {
		super(message);
	}
	
	public final CardExceptionType getExceptionType() {
		return exceptionType;
	}
}
